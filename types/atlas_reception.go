package types

// RepoRun type is the expected structure of a repo run task
// to be received
type RepoRun struct {
	RunID   string `json:"run_id"`
	VCSMeta struct {
		Name         string `json:"name"`
		Provider     string `json:"provider"`
		RepoURL      string `json:"repo_url"`
		Owner        string `json:"owner"`
		Repository   string `json:"repository"`
		CheckoutHash string `json:"checkout_hash"`
	} `json:"vcs_meta"`
}

// AnalysisRun type is the expected structure of a analysis run task
// to be received
type AnalysisRun struct {
	RunID   string `json:"run_id"`
	VCSMeta struct {
		Name         string `json:"name"`
		Provider     string `json:"provider"`
		RepoURL      string `json:"repo_url"`
		Owner        string `json:"owner"`
		Repository   string `json:"repository"`
		DefaultHash  string `json:"default_hash"`
		CheckoutHash string `json:"checkout_hash"`
	} `json:"vcs_meta"`
	AnalyzerMeta struct {
		Shortcode   string `json:"name"`
		Command     string `json:"command"`
		Version     string `json:"version"`
		CPULimit    string `json:"cpu_limit"`
		MemoryLimit string `json:"memory_limit"`
	}
	Processors struct {
		SourceCodeLoad bool `json:"source_code_load"`
		Beacon         bool `json:"beacon"`
	} `json:"processors"`
	Filters []struct {
		Name       string   `json:"name"`
		Expression string   `json:"expression"`
		Variables  []string `json:"variables"`
		Mutations  []string `json:"mutations"`
	} `json:"filters"`
}

// Beacon type is the expected structure of a beacon task
// to be received
type BeaconRun struct {
	RunID        string `json:"run_id"`
	RepositoryID int64  `json:"repository_id"`
}
