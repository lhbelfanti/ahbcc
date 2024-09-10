package criteria

import (
	"net/http"
	"strconv"

	"ahbcc/internal/log"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/enqueue/v1
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

		err = enqueueCriteria(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToEnqueueCriteria, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "Criteria successfully sent to enqueue")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria successfully sent to enqueue"))
	}
}
