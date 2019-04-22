package types

// RepoRun type is the expected structure of a repo run task
// to be received
type RepoRun struct {
	RunID     string `json:"run_id"`
	RunSerial string `json:"run_serial"`
	VCSMeta   struct {
		RemoteURL   string `json:"remote_url"`
		CheckoutOID string `json:"checkout_oid"`
	} `json:"vcs_meta"`
}

// AnalysisRun type is the expected structure of a analysis run task
// to be received
type AnalysisRun struct {
	RunID           string `json:"run_id"`
	RunSerial       string `json:"run_serial"`
	DSConfigUpdated bool   `json:"ds_config_updated"`
	VCSMeta         struct {
		RemoteURL   string `json:"remote_url"`
		BaseBranch  string `json:"base_branch"`
		BaseOID     string `json:"base_oid"`
		CheckoutOID string `json:"checkout_oid"`
	} `json:"vcs_meta"`
	Checks []struct {
		CheckSeq     string  `json:"check_seq"`
		Artifacts    []int64 `json:"artifacts"`
		AnalyzerMeta struct {
			Shortcode   string `json:"name"`
			Command     string `json:"command"`
			Version     string `json:"version"`
			CPULimit    string `json:"cpu_limit"`
			MemoryLimit string `json:"memory_limit"`
		} `json:"analyzer_meta"`
		Processors []string `json:"processors"`
	} `json:"checks"`
}

// CancelCheckRun type is the expected structure of a check cancellation
// task to be recieved
type CancelCheckRun struct {
	RunID        string `json:"run_id"`
	RunSerial    string `json:"run_serial"`
	AnalysisMeta struct {
		RunID     string `json:"run_id"`
		RunSerial string `json:"run_serial"`
		CheckSeq  string `json:"check_seq"`
	} `json:"analysis_meta"`
}

// Beacon type is the expected structure of a beacon task
// to be received
type BeaconRun struct {
	RunID        string `json:"run_id"`
	RepositoryID int64  `json:"repository_id"`
}
