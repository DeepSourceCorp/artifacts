package types

type Change struct {
	BeforeHTML string `json:"before_html"`
	AfterHTML  string `json:"after_html"`
	Changeset  string `json:"changeset"`
	Identifier string `json:"identifier"`
}

type Patch struct {
	Filename string   `json:"filename"`
	Changes  []Change `json:"changes"`
	Action   string   `json:"action"`
}

type FixedIssue struct {
	IssueCode string `json:"issue_code"`
	Count     int    `json:"count"`
}

type IssuesFixed struct {
	Filename    string       `json:"filename"`
	FixedIssues []FixedIssue `json:"fixed_issues"`
}

type AutofixReport struct {
	CodeDir      string        `json:"code_directory,omitempty"`
	ChangedFiles []string      `json:"changed_files,omitempty"`
	IssuesFixed  []IssuesFixed `json:"issues_fixed"`
	Metrics      []Metric      `json:"metrics,omitempty"`
	Patches      []Patch       `json:"patches"`
	Errors       []Error       `json:"errors"`
	ExtraData    interface{}   `json:"extra_data"`
}

type AutofixResult struct {
	RunID    string        `json:"run_id"`
	Status   Status        `json:"status"`
	CheckSeq string        `json:"check_seq"`
	Report   AutofixReport `json:"report"`
}

type AutofixResultCeleryTask struct {
	ID      string        `json:"id"`
	Task    string        `json:"task"`
	KWArgs  AutofixResult `json:"kwargs"`
	Retries int           `json:"retries"`
}

// Issues to be autofixed
//
//proteus:generate
type AutofixIssue struct {
	IssueCode   string          `json:"issue_code"`
	Occurrences []IssueLocation `json:"occurrences"`
}

//proteus:generate
type AutofixConfig struct {
	Issues []AutofixIssue `json:"issues"`
	Meta   interface{}    `json:"meta"`
}

//proteus:generate
type MarvinAutofixConfig struct {
	RunID             string `toml:"runID"`
	AnalyzerShortcode string `toml:"analyzerShortcode"`
	AutofixerCommand  string `toml:"autofixerCommand"`
	CheckoutOID       string `toml:"checkoutOID"`
	AutofixIssues     string `toml:"autofix_issues"`
}

// Transformers types.
type TransformerReport struct {
	CodeDir      string   `json:"code_directory,omitempty"`
	ChangedFiles []string `json:"changed_files,omitempty"`
	Errors       []Error  `json:"errors"`
	Patches      []Patch  `json:"patches"`
}

type TransformerResult struct {
	RunID  string            `json:"run_id"`
	Status Status            `json:"status"`
	Report TransformerReport `json:"report"`
}

type TransformerResultCeleryTask struct {
	ID      string            `json:"id"`
	Task    string            `json:"task"`
	KWArgs  TransformerResult `json:"kwargs"`
	Retries int               `json:"retries"`
}

type TransformerCommitData struct {
	Branch  string `toml:"branch"`
	Author  string `toml:"author"`
	Email   string `toml:"email"`
	Message string `toml:"message"`
}

type MarvinTransformerConfig struct {
	RunID                 string                `toml:"runID"`
	BaseOID               string                `toml:"baseOID"`
	CheckoutOID           string                `toml:"checkoutOID"`
	IsTriggeredByRunner   bool                  `toml:"is_triggered_by_runner"`
	TransformerCommand    string                `toml:"transformerCommand"`
	TransformerTools      []string              `toml:"transformerTools"`
	DSConfigUpdated       bool                  `toml:"dsConfigUpdated"`
	TransformerCommitMeta TransformerCommitData `toml:"transformer_commit_meta"`
}

type TransformerConfig struct {
	ExcludePatterns []string `json:"exclude_patterns"`
	ExcludeFiles    []string `json:"exclude_files"`
	Files           []string `json:"files"`
	Tools           []string `json:"tools"`
}
