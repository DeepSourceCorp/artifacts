package types

type MarvinConfig struct {
	MachineryTaskID   string   `toml:"taskID"`
	RunID             string   `toml:"runID"`
	RunType           string   `toml:"runType"`
	RMQRoutingKey     string   `toml:"rmqRoutingKey"`
	CheckSeq          string   `toml:"checkSeq"`
	AnalyzerShortcode string   `toml:"analyzerShortcode"`
	AnalyzerCommand   string   `toml:"analyzerCommand"`
	DefaultHash       string   `toml:"defaultHash"`
	CheckoutHash      string   `toml:"checkoutHash"`
	DSConfigUpdated   bool     `toml:"dsConfigUpdated"`
	Processors        []string `toml:"processors"`
}

type AnalysisConfig struct {
	Files           []string    `json:"files"`
	ExcludePatterns []string    `json:"exclude_patterns"`
	AnalyzerMeta    interface{} `json:"analyzer_meta"`
}
