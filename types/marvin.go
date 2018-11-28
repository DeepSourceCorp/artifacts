package types

type MarvinConfig struct {
	MachineryTaskID string `toml:"taskID"`
	RunID           string `toml:"runID"`
	RunType         string `toml:"runType"`
	RMQURL          string `toml:"rmqURL"`
	RMQExchange     string `toml:"rmqExchange"`
	RMQRoutingKey   string `toml:"rmqRoutingKey"`
	CheckSeq        string `toml:"checkSeq"`
	TaskType        string `toml:"taskType"`
	TaskShortcode   string `toml:"taskShortcode"`
	TaskCommand     string `toml:"taskCommand"`
	DefaultHash     string `toml:"defaultHash"`
	CheckoutHash    string `toml:"checkoutHash"`
}

type AnalysisConfig struct {
	Files           []string `json:"files"`
	ExcludePatterns []string `json:"exclude_patterns"`
	Processors      []string `json:"processors"`
}
