package types

// RepoRun type is the expected structure of a repo run task
// to be received
//proteus:generate
type RepoRunVCSMeta struct {
	RemoteURL   string `json:"remote_url"`
	CheckoutOID string `json:"checkout_oid"`
}

//proteus:generate
type RepoRun struct {
	RunID     string         `json:"run_id"`
	RunSerial string         `json:"run_serial"`
	VCSMeta   RepoRunVCSMeta `json:"vcs_meta"`
}

//proteus:generate
type Artifact struct {
	Key      string            `json:"key"`
	URL      string            `json:"url"`
	Metadata map[string]string `json:"metadata"`
}

// AnalysisRun type is the expected structure of a analysis run task
// to be received
//proteus:generate
type AnalysisRunVCSMeta struct {
	RemoteURL       string `json:"remote_url"`
	BaseBranch      string `json:"base_branch"`
	BaseOID         string `json:"base_oid"`
	CheckoutOID     string `json:"checkout_oid"`
	CloneSubmodules bool   `json:"clone_submodules"`
}

//proteus:generate
type SSH struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

//proteus:generate
type Keys struct {
	SSH SSH `json:"ssh,omitempty"`
}

//proteus:generate
type AnalyzerMeta struct {
	Shortcode    string `json:"name"`
	AnalyzerType string `json:"analyzer_type"`
	Command      string `json:"command"`
	Version      string `json:"version"`
	CPULimit     string `json:"cpu_limit"`
	MemoryLimit  string `json:"memory_limit"`
}

//proteus:generate
type Check struct {
	CheckSeq        string           `json:"check_seq"`
	Artifacts       []Artifact       `json:"artifacts"`
	AnalyzerMeta    AnalyzerMeta     `json:"analyzer_meta"`
	Processors      []string         `json:"processors"`
	DiffMetaCommits []DiffMetaCommit `json:"diff_meta_commits"`
}

type DiffMetaCommit struct {
	CommitOID string   `json:"commit_oid" toml:"commitOID"`
	Paths     []string `json:"paths" toml:"paths"`
}

//proteus:generate
type AnalysisRun struct {
	RunID           string             `json:"run_id"`
	RunSerial       string             `json:"run_serial"`
	Config          DSConfig           `json:"config"`
	DSConfigUpdated bool               `json:"ds_config_updated"`
	VCSMeta         AnalysisRunVCSMeta `json:"vcs_meta"`
	Keys            Keys               `json:"keys"`
	Checks          []Check            `json:"checks"`
	Meta            map[string]string  `json:"_meta"`
}

//proteus:generate
type AutofixVCSMeta struct {
	RemoteURL       string `json:"remote_url"`
	BaseBranch      string `json:"base_branch"`
	CheckoutOID     string `json:"checkout_oid"`
	CloneSubmodules bool   `json:"clone_submodules"`
}

//proteus:generate
type AutofixMeta struct {
	Shortcode   string `json:"name"`
	Command     string `json:"command"`
	Version     string `json:"version"`
	CPULimit    string `json:"cpu_limit"`
	MemoryLimit string `json:"memory_limit"`
}

//proteus:generate
type Autofixer struct {
	AutofixMeta AutofixMeta    `json:"autofix_meta"`
	Autofixes   []AutofixIssue `json:"autofixes"`
}

//proteus:generate
type AutofixRun struct {
	RunID     string            `json:"run_id"`
	RunSerial string            `json:"run_serial"`
	Config    DSConfig          `json:"config"`
	VCSMeta   AutofixVCSMeta    `json:"vcs_meta"`
	Keys      Keys              `json:"keys"`
	Autofixer Autofixer         `json:"autofixer"`
	Meta      map[string]string `json:"_meta"`
}

type TransformerVCSMeta struct {
	RemoteURL       string `json:"remote_url"`
	BaseBranch      string `json:"base_branch"`
	BaseOID         string `json:"base_oid"`
	CheckoutOID     string `json:"checkout_oid"`
	CloneSubmodules bool   `json:"clone_submodules"`
}

type TransformerMeta struct {
	Version     string `json:"version"`
	CPULimit    string `json:"cpu_limit"`
	MemoryLimit string `json:"memory_limit"`
}

type TransformerInfo struct {
	Command string          `json:"command"`
	Tools   []string        `json:"tools"`
	Meta    TransformerMeta `json:"meta"`
}

type TransformerRun struct {
	RunID           string             `json:"run_id"`
	RunSerial       string             `json:"run_serial"`
	Config          DSConfig           `json:"config"`
	VCSMeta         TransformerVCSMeta `json:"vcs_meta"`
	DSConfigUpdated bool               `json:"ds_config_updated"`
	Transformer     TransformerInfo    `json:"transformer"`
	Meta            map[string]string  `json:"_meta"`
}

// CancelCheckRun type is the expected structure of a check cancellation
// task to be recieved
//proteus:generate
type CancelCheckAnalysisMeta struct {
	RunID     string `json:"run_id"`
	RunSerial string `json:"run_serial"`
	CheckSeq  string `json:"check_seq"`
}

//proteus:generate
type CancelCheckRun struct {
	AnalysisMeta CancelCheckAnalysisMeta `json:"analysis_meta"`
	RunID        string                  `json:"run_id"`
	RunSerial    string                  `json:"run_serial"`
}

// Beacon type is the expected structure of a beacon task
// to be received
//proteus:generate
type BeaconRun struct {
	RunID        string `json:"run_id"`
	RepositoryID int64  `json:"repository_id"`
}
