package rules

import (
	"ahbcc/internal/http/response"
	"ahbcc/internal/log"
	"encoding/json"
	"net/http"
	"slices"
)

// InsertRulesHandlerV1 HTTP Handler of the endpoint /corpus/cleaning-rules/v1
func InsertRulesHandlerV1(insertRules Insert) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body []DTO
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("rules", body))

		err = validateBody(body)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		err = insertRules(ctx, body)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToInsertRulesIntoDatabase, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Rules successfully inserted", nil, nil)
	}
}

// validateBody validates the rule type of the given rules
func validateBody(body []DTO) error {
	for _, rule := range body {
		if !slices.Contains(validRuleTypes, rule.RuleType) {
			return WrongRuleType
		}
	}

	return nil
}
