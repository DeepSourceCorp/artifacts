package types

type MarvinConfig struct {
	RunID             string   `toml:"runID"`
	CheckSeq          string   `toml:"checkSeq"`
	AnalyzerShortcode string   `toml:"analyzerShortcode"`
	AnalyzerCommand   string   `toml:"analyzerCommand"`
	DefaultOID        string   `toml:"defaultOID"`
	CheckoutOID       string   `toml:"checkoutOID"`
	DSConfigUpdated   bool     `toml:"dsConfigUpdated"`
	Processors        []string `toml:"processors"`
}

type AnalysisConfig struct {
	Files           []string    `json:"files"`
	ExcludePatterns []string    `json:"exclude_patterns"`
	AnalyzerMeta    interface{} `json:"analyzer_meta"`
}
