package corpus

import (
	"net/http"

	"ahbcc/internal/http/response"
)

// CreateCorpusHandlerV1 HTTP Handler of the endpoint /corpus/v1
func CreateCorpusHandlerV1(createCorpus Create) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := createCorpus(ctx)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToCreateCorpus, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Corpus successfully created", nil, nil)
	}
}
