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
				return
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

// UpdateExecutionHandlerV1 HTTP Handler of the endpoint /criteria/executions/{execution_id}/v1
func UpdateExecutionHandlerV1(updateExecution UpdateExecution) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		executionIDParam := r.PathValue("execution_id")
		executionID, err := strconv.Atoi(executionIDParam)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidURLParameter, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("execution_id", executionIDParam))

		var execution ExecutionDTO
		err = json.NewDecoder(r.Body).Decode(&execution)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("execution", execution))

		err = updateExecution(ctx, executionID, execution.Status)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToExecuteUpdateCriteriaExecution, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "Criteria execution successfully updated")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria execution successfully updated"))
	}
}

// InsertExecutionDayHandlerV1 HTTP Handler of the endpoint /criteria/executions/{execution_id}/day/v1
func InsertExecutionDayHandlerV1(insertExecutionDay InsertExecutionDay) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		executionIDParam := r.PathValue("execution_id")
		executionID, err := strconv.Atoi(executionIDParam)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidURLParameter, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("execution_id", executionIDParam))

		var executionDay ExecutionDayDTO
		err = json.NewDecoder(r.Body).Decode(&executionDay)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("execution_day", executionDay))

		executionDay.SearchCriteriaExecutionID = executionID

		err = insertExecutionDay(ctx, executionDay)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToExecuteInsertCriteriaExecution, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "Criteria execution day successfully inserted")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria execution day successfully inserted"))
	}
}
