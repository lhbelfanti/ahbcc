package criteria

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/lhbelfanti/corpus-creator/internal/http/response"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/{criteria_id}/enqueue/v1
func EnqueueHandlerV1(enqueueCriteria Enqueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		criteriaIDParam := r.PathValue("criteria_id")
		criteriaID, err := strconv.Atoi(criteriaIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("criteria_id", criteriaIDParam))

		forcedQueryParamStr := r.URL.Query().Get("forced")
		forcedQueryParam, err := strconv.ParseBool(forcedQueryParamStr)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("forced", forcedQueryParamStr))

		err = enqueueCriteria(ctx, criteriaID, forcedQueryParam)
		if err != nil {
			switch {
			case errors.Is(err, AnExecutionOfThisCriteriaIDIsAlreadyEnqueued):
				response.Send(ctx, w, http.StatusConflict, ExecutionWithSameCriteriaIDAlreadyEnqueued, nil, err)
				return
			default:
				response.Send(ctx, w, http.StatusInternalServerError, FailedToEnqueueCriteria, nil, err)
				return
			}
		}

		response.Send(ctx, w, http.StatusOK, "Criteria successfully sent to enqueue", nil, nil)
	}
}

// InitHandlerV1 HTTP Handler of the endpoint /criteria/init/v1
func InitHandlerV1(init Init) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := init(ctx)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteInitCriteria, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria successfully initialized and enqueued", nil, nil)
	}
}

// InformationHandlerV1 HTTP Handler of the endpoint /criteria/v1
func InformationHandlerV1(information Information) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			response.Send(ctx, w, http.StatusUnauthorized, AuthorizationTokenRequired, nil, AuthorizationTokenIsRequired)
			return
		}

		criteriaInformation, err := information(ctx, token)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteCriteriaInformation, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria successfully obtained", criteriaInformation, nil)
	}
}

// SummarizedInformationHandlerV1 HTTP Handler of the endpoint /criteria/{criteria_id}/summarize/v1
func SummarizedInformationHandlerV1(summarizedInformation SummarizedInformation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var year, month int

		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			response.Send(ctx, w, http.StatusUnauthorized, AuthorizationTokenRequired, nil, AuthorizationTokenIsRequired)
			return
		}

		criteriaIDParam := r.PathValue("criteria_id")
		criteriaID, err := strconv.Atoi(criteriaIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("criteria_id", criteriaIDParam))

		yearQueryParamStr := r.URL.Query().Get("year")
		if yearQueryParamStr != "" {
			year, err = strconv.Atoi(yearQueryParamStr)
			if err != nil {
				response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
				return
			}
			ctx = log.With(ctx, log.Param("year", yearQueryParamStr))
		}

		monthQueryParamStr := r.URL.Query().Get("month")
		if monthQueryParamStr != "" {
			month, err = strconv.Atoi(monthQueryParamStr)
			if err != nil {
				response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
				return
			}
			ctx = log.With(ctx, log.Param("month", monthQueryParamStr))
		}

		criteriaInformation, err := summarizedInformation(ctx, token, criteriaID, year, month)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteCriteriaSummarizedInformation, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria successfully obtained", criteriaInformation, nil)
	}
}
