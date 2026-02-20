package types

// VendorAnalysisRunPositionDetail represents a line/column location in a file.
type VendorAnalysisRunPositionDetail struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// VendorAnalysisRunPosition represents the begin/end span of a finding in a file.
type VendorAnalysisRunPosition struct {
	Begin VendorAnalysisRunPositionDetail `json:"begin"`
	End   VendorAnalysisRunPositionDetail `json:"end"`
}

// VendorAnalysisRunSecretProvider identifies the provider of a detected secret.
type VendorAnalysisRunSecretProvider struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// VendorAnalysisRunSecretValidation holds the validation request/response for a secret.
type VendorAnalysisRunSecretValidation struct {
	Request          interface{} `json:"request"`
	ExpectedResponse interface{} `json:"expected_response"`
}

// VendorAnalysisRunIssue represents a single detected issue.
type VendorAnalysisRunIssue struct {
	Object           string                             `json:"object"`
	File             string                             `json:"file"`
	Position         VendorAnalysisRunPosition          `json:"position"`
	Explanation      *string                            `json:"explanation,omitempty"`
	Category         string                             `json:"category"`
	Provider         *VendorAnalysisRunSecretProvider   `json:"provider,omitempty"`
	Value            *string                            `json:"value,omitempty"`
	SecretValidation *VendorAnalysisRunSecretValidation `json:"secret_validation,omitempty"`
	Language         *string                            `json:"language,omitempty"`
}

// VendorAnalysisRunFix represents a single fix produced for a detected issue.
type VendorAnalysisRunFix struct {
	Object      string  `json:"object"`
	Category    string  `json:"category"`
	Patch       string  `json:"patch"`
	Explanation string  `json:"explanation"`
	Language    *string `json:"language,omitempty"`
}

// VendorAnalysisRunDetectionResult aggregates all detection findings for a run.
type VendorAnalysisRunDetectionResult struct {
	Object                   string         `json:"object"`
	Issues                   []VendorAnalysisRunIssue `json:"issues"`
	IssuesDetectedByCategory map[string]int `json:"issues_detected_by_category"`
	IssuesDetectedByLanguage map[string]int `json:"issues_detected_by_language"`
	IssuesDetectedCount      int            `json:"issues_detected_count"`
}

// VendorAnalysisRunFixResult aggregates all fixes produced for a run.
type VendorAnalysisRunFixResult struct {
	Object                string                 `json:"object"`
	Patch                 string                 `json:"patch"`
	IssuesFixedCount      int                    `json:"issues_fixed_count"`
	Fixes                 []VendorAnalysisRunFix `json:"fixes"`
	IssuesFixedByLanguage map[string]int         `json:"issues_fixed_by_language"`
	IssuesFixedByCategory map[string]int         `json:"issues_fixed_by_category"`
}

// VendorAnalysisRunResult is the top-level result published by the vendor analysis
// pipeline upon completion. It is the Go equivalent of
// enki_v2/vendor_service/models.py:AnalysisReport.
type VendorAnalysisRunResult struct {
	DetectionResult VendorAnalysisRunDetectionResult `json:"detection_result"`
	FixResult       VendorAnalysisRunFixResult       `json:"fix_result"`
}

// VendorAnalysisRunResultCeleryTask is the Celery task envelope used to deliver a
// VendorAnalysisRunResult back to asgard via RabbitMQ.
type VendorAnalysisRunResultCeleryTask struct {
	ID      string                  `json:"id"`
	Task    string                  `json:"task"`
	KWArgs  VendorAnalysisRunResult `json:"kwargs"`
	Retries int                     `json:"retries"`
}
