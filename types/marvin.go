package types

type MarvinConfig struct {
	RunID             string   `toml:"runID"`
	CheckSeq          string   `toml:"checkSeq"`
	AnalyzerShortcode string   `toml:"analyzerShortcode"`
	AnalyzerCommand   string   `toml:"analyzerCommand"`
	BaseOID           string   `toml:"baseOID"`
	CheckoutOID       string   `toml:"checkoutOID"`
	DSConfigUpdated   bool     `toml:"dsConfigUpdated"`
	Processors        []string `toml:"processors"`
}

type AnalysisConfig struct {
	Files           []string    `json:"files"`
	ExcludePatterns []string    `json:"exclude_patterns"`
	ExcludeFiles    []string    `json:"exclude_files"`
	TestFiles       []string    `json:"test_files"`
	TestPatterns    []string    `json:"test_patterns"`
	AnalyzerMeta    interface{} `json:"analyzer_meta"`
}

type AnalysisStateInfo struct {
	IfAllFiles bool `json:"if_all_files"`
}
