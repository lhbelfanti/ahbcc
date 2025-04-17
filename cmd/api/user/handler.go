package user

import (
	"net/http"
	"strconv"

	"ahbcc/internal/http/response"
	"ahbcc/internal/log"
)

// InformationV1 HTTP Handler of the endpoint /users/{user_id}/criteria/v1
func InformationV1(information Information) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userIDParam := r.PathValue("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("user_id", userIDParam))

		criteriaInformation, err := information(ctx, userID)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExecuteCriteriaInformation, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria successfully obtained", criteriaInformation, nil)
	}
}
