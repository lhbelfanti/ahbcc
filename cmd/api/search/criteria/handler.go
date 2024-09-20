package criteria

import (
	"errors"
	"net/http"
	"strconv"

	"ahbcc/internal/log"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/{criteria_id}/enqueue/v1
func EnqueueHandlerV1(enqueueCriteria Enqueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		criteriaIDParam := r.PathValue("criteria_id")
		criteriaID, err := strconv.Atoi(criteriaIDParam)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidURLParameter, http.StatusBadRequest)
			return
		}

		forcedQueryParamStr := r.URL.Query().Get("forced")
		forcedQueryParam, err := strconv.ParseBool(forcedQueryParamStr)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidQueryParameterFormat, http.StatusBadRequest)
			return
		}

		err = enqueueCriteria(ctx, criteriaID, forcedQueryParam)
		if err != nil {
			switch {
			case errors.Is(err, AnExecutionOfThisCriteriaIDIsAlreadyEnqueued):
				log.Error(ctx, err.Error())
				http.Error(w, ExecutionWithSameCriteriaIDAlreadyEnqueued, http.StatusConflict)
			default:
				log.Error(ctx, err.Error())
				http.Error(w, FailedToEnqueueCriteria, http.StatusInternalServerError)
				return
			}
		}

		log.Info(ctx, "Criteria successfully sent to enqueue")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria successfully sent to enqueue"))
	}
}

// InitHandlerV1 HTTP Handler of the endpoint /criteria/init/v1
func InitHandlerV1(init Init) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := init(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToExecuteInitCriteria, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "Criteria successfully initialized and enqueued")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria successfully initialized and enqueued"))
	}
}
