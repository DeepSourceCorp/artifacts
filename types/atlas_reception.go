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
	ParallelCount int `json:"parallel_count"`
	Checks        []struct {
		CheckSeq string `json:"check_seq"`
		Config   struct {
			SourceCodeLoad bool `json:"source_code_load"`
		} `json:"config"`
		Tasks []struct {
			TaskType string `json:"task_type"`
			TaskMeta struct {
				Shortcode   string `json:"name"`
				Command     string `json:"command"`
				Version     string `json:"version"`
				CPULimit    string `json:"cpu_limit"`
				MemoryLimit string `json:"memory_limit"`
			} `json:"task_meta"`
		} `json:"tasks"`
	} `json:"checks"`
}
