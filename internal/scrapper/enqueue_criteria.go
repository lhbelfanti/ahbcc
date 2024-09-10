package scrapper

import (
	"context"
	"fmt"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/http"
	"ahbcc/internal/log"
)

// EnqueueCriteria calls the endpoint to enqueue a criteria seeking by the criteria ID
type EnqueueCriteria func(ctx context.Context, criteriaID int) error

// MakeEnqueueCriteria creates a new EnqueueCriteria
func MakeEnqueueCriteria(httpClient http.Client, domain string, selectCriteriaByID criteria.SelectByID) EnqueueCriteria {
	url := domain + "/tweets/v1"

	return func(ctx context.Context, criteriaID int) error {
		criteriaToEnqueue, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectCriteriaByID
		}

		resp, err := httpClient.NewRequest(ctx, "POST", url, criteriaToEnqueue)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}

		log.Info(ctx, fmt.Sprintf("Enqueue search criteria endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}
