package types

// AIAutofixIssueResult represents autofix patches for a single issue, linked by UUID.
type AIAutofixIssueResult struct {
	IssueUUID           string   `json:"issue_uuid"`
	FixPatches          []string `json:"fix_patches,omitempty"`
	FixTitle            string   `json:"fix_title,omitempty"`
	FixExplanation      string   `json:"fix_explanation,omitempty"`
	FixReplacementTexts []string `json:"fix_replacement_texts,omitempty"`
}

// AIAutofixResults is the top-level structure of autofix_results.json written by enki.
type AIAutofixResults struct {
	Issues []AIAutofixIssueResult `json:"issues"`
}

// AIAutofixResult is the payload published to the backend via RMQ.
type AIAutofixResult struct {
	RunID        string           `json:"run_id"`
	CheckSeq     string           `json:"check_seq"`
	ReportObject string           `json:"report_object"`
	Results      AIAutofixResults `json:"results"`
}

// AIAutofixResultCeleryTask wraps AIAutofixResult for Celery task publishing.
type AIAutofixResultCeleryTask struct {
	ID      string          `json:"id"`
	Task    string          `json:"task"`
	KWArgs  AIAutofixResult `json:"kwargs"`
	Retries int             `json:"retries"`
}
