package criteria

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"ahbcc/internal/http/response"
	"ahbcc/internal/log"
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

// GetExecutionByIDHandlerV1 HTTP Handler of the endpoint /criteria/executions/{execution_id}/v1
func GetExecutionByIDHandlerV1(selectExecutionByID SelectExecutionByID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		executionIDParam := r.PathValue("execution_id")
		executionID, err := strconv.Atoi(executionIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("execution_id", executionIDParam))

		executions, err := selectExecutionByID(ctx, executionID)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteGetExecutionsByStatuses, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Execution successfully obtained", executions, nil)
	}
}

// UpdateExecutionHandlerV1 HTTP Handler of the endpoint /criteria/executions/{execution_id}/v1
func UpdateExecutionHandlerV1(updateExecution UpdateExecution) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		executionIDParam := r.PathValue("execution_id")
		executionID, err := strconv.Atoi(executionIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("execution_id", executionIDParam))

		var execution ExecutionDTO
		err = json.NewDecoder(r.Body).Decode(&execution)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("execution", execution))

		err = updateExecution(ctx, executionID, execution.Status)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteUpdateCriteriaExecution, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria execution successfully updated", nil, nil)
	}
}

// CreateExecutionDayHandlerV1 HTTP Handler of the endpoint /criteria/executions/{execution_id}/day/v1
func CreateExecutionDayHandlerV1(insertExecutionDay InsertExecutionDay) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		executionIDParam := r.PathValue("execution_id")
		executionID, err := strconv.Atoi(executionIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("execution_id", executionIDParam))

		var executionDay ExecutionDayDTO
		err = json.NewDecoder(r.Body).Decode(&executionDay)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("execution_day", executionDay))

		executionDay.SearchCriteriaExecutionID = executionID

		err = insertExecutionDay(ctx, executionDay)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteInsertCriteriaExecution, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria execution day successfully inserted", nil, nil)
	}
}
