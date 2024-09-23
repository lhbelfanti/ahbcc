package scrapper

import (
	"context"
	"fmt"

	"ahbcc/internal/http"
	"ahbcc/internal/log"
)

// EnqueueCriteria calls the endpoint to enqueue a criteria seeking by the criteria ID
type EnqueueCriteria func(ctx context.Context, body CriteriaDTO) error

// MakeEnqueueCriteria creates a new EnqueueCriteria
func MakeEnqueueCriteria(httpClient http.Client, domain string) EnqueueCriteria {
	url := domain + "/criteria/enqueue/v1"

	return func(ctx context.Context, criteria CriteriaDTO) error {
		body := EnqueueCriteriaMessageDTO{Message: criteria}
		resp, err := httpClient.NewRequest(ctx, "POST", url, body)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}
		ctx = log.With(ctx, log.Param("body", body))

		log.Info(ctx, fmt.Sprintf("Enqueue search criteria endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}
