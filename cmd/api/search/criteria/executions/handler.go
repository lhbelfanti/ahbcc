package executions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lhbelfanti/corpus-creator/internal/http/response"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// GetExecutionByIDHandlerV1 HTTP Handler of the endpoint /criteria-executions/{execution_id}/v1
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

		execution, err := selectExecutionByID(ctx, executionID)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteGetExecutionsByStatuses, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Execution successfully obtained", execution, nil)
	}
}

// UpdateExecutionHandlerV1 HTTP Handler of the endpoint /criteria-executions/{execution_id}/v1
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

// CreateExecutionDayHandlerV1 HTTP Handler of the endpoint /criteria-executions/{execution_id}/day/v1
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

// SummarizeHandlerV1 HTTP Handler of the endpoint /criteria-executions/summarize/v1
func SummarizeHandlerV1(summarize Summarize) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := summarize(ctx)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteSummarize, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria executions summarization successfully run", nil, nil)
	}
}
