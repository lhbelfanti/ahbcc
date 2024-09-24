package criteria

import (
	"encoding/json"
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
		ctx = log.With(ctx, log.Param("criteria_id", criteriaIDParam))

		forcedQueryParamStr := r.URL.Query().Get("forced")
		forcedQueryParam, err := strconv.ParseBool(forcedQueryParamStr)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidQueryParameterFormat, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("forced", forcedQueryParamStr))

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

// InsertExecutionHandlerV1 HTTP Handler of the endpoint /criteria/{criteria_id}/executions/{execution_id}/v1
func InsertExecutionHandlerV1(insertExecution InsertExecution) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		criteriaIDParam := r.PathValue("criteria_id")
		criteriaID, err := strconv.Atoi(criteriaIDParam)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidURLParameter, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("criteria_id", criteriaIDParam))

		executionID, err := insertExecution(ctx, criteriaID, true)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToExecuteInsertCriteriaExecution, http.StatusInternalServerError)
			return
		}

		response := InsertExecutionHandlerV1Response{
			Message:     "Criteria execution successfully inserted",
			ExecutionID: executionID,
		}

		log.Info(ctx, "Criteria execution successfully inserted")
		w.Header().Set("Content-Type", "application/json")
		
		err = json.NewEncoder(w).Encode(response) // w.WriteHeader(http.StatusOK) is implicit set inside the encoder
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(ctx, err.Error())
			http.Error(w, FailedToEncodeInsertCriteriaExecutionResponse, http.StatusInternalServerError)
			return
		}
	}
}
