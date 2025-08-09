package corpus

// ExportResult represents the information needed by the handler to export the corpus
type ExportResult struct {
	Data        []byte
	ContentType string
	Filename    string
}
