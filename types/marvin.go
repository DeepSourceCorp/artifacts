package types

//proteus:generate
type MarvinAnalysisConfig struct {
	RunID             string           `toml:"runID"`
	CheckSeq          string           `toml:"checkSeq"`
	AnalyzerShortcode string           `toml:"analyzerShortcode"`
	AnalyzerCommand   string           `toml:"analyzerCommand"`
	BaseOID           string           `toml:"baseOID"`
	CheckoutOID       string           `toml:"checkoutOID"`
	DSConfigUpdated   bool             `toml:"dsConfigUpdated"`
	Processors        []string         `toml:"processors"`
	DiffMetaCommits   []DiffMetaCommit `toml:"diffMetaCommits"`
}

//proteus:generate
type MarvinAutofixConfig struct {
	RunID             string `toml:"runID"`
	AnalyzerShortcode string `toml:"analyzerShortcode"`
	AutofixerCommand  string `toml:"autofixerCommand"`
	CheckoutOID       string `toml:"checkoutOID"`
	AutofixIssues     string `toml:"autofix_issues"`
}

type MarvinTransformerConfig struct {
	RunID              string   `toml:"runID"`
	BaseOID            string   `toml:"baseOID"`
	CheckoutOID        string   `toml:"checkoutOID"`
	TransformerCommand string   `toml:"transformerCommand"`
	TransformerTools   []string `toml:"transformerTools"`
	DSConfigUpdated    bool     `toml:"dsConfigUpdated"`
}

//proteus:generate
type AnalysisConfig struct {
	Files           []string    `json:"files"`
	ExcludePatterns []string    `json:"exclude_patterns"`
	ExcludeFiles    []string    `json:"exclude_files"`
	TestFiles       []string    `json:"test_files"`
	TestPatterns    []string    `json:"test_patterns"`
	AnalyzerMeta    interface{} `json:"analyzer_meta"`
}

//proteus:generate
type AnalysisStateInfo struct {
	IfAllFiles bool `json:"if_all_files"`
}

// Issues to be autofixed
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

type TransformerConfig struct {
	ExcludePatterns []string `json:"exclude_patterns"`
	ExcludeFiles    []string `json:"exclude_files"`
	Files           []string `json:"files"`
	Tools           []string `json:"tools"`
}
