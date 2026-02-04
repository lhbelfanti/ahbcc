package corpus

import (
	"net/http"

	"ahbcc/internal/http/response"
)

// CreateCorpusHandlerV1 HTTP Handler of the endpoint /corpus/v1
func CreateCorpusHandlerV1(createCorpus Create) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var perfectBalancedCorpus bool
		perfectBalancedCorpusQueryString := r.URL.Query().Get("perfectBalancedCorpus")
		if perfectBalancedCorpusQueryString == "true" {
			perfectBalancedCorpus = true
		}

		err := createCorpus(ctx, perfectBalancedCorpus)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToCreateCorpus, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Corpus successfully created", nil, nil)
	}
}

// ExportCorpusHandlerV1 HTTP Handler of the endpoint /corpus/v1
func ExportCorpusHandlerV1(exportCorpus ExportCorpus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		format := r.URL.Query().Get("format")
		if format == "" {
			format = JSONFormat
		}

		result, err := exportCorpus(ctx, format)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToExportCorpus, nil, err)
			return
		}

		w.Header().Set("Content-Type", result.ContentType)
		w.Header().Set("Content-Disposition", "attachment; filename="+result.Filename)
		w.Write(result.Data)
	}
}
