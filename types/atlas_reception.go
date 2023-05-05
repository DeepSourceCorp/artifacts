package types

// RepoRun type is the expected structure of a repo run task
// to be received
//
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
//
//proteus:generate
type AnalysisRunVCSMeta struct {
	RemoteURL                  string `json:"remote_url"`
	BaseBranch                 string `json:"base_branch"`
	BaseOID                    string `json:"base_oid"`
	CheckoutOID                string `json:"checkout_oid"`
	IsForDefaultAnalysisBranch bool   `json:"is_for_default_analysis_branch"`
	CloneSubmodules            bool   `json:"clone_submodules"`
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
	ImagePath    string `json:"image_path"`
	AnalyzerType string `json:"analyzer_type"`
	Command      string `json:"command"`
	Version      string `json:"version"`
	CPULimit     string `json:"cpu_limit"`
	MemoryLimit  string `json:"memory_limit"`
	CacheVersion int    `json:"cache_version"`
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
	Shortcode    string `json:"name"`
	Command      string `json:"command"`
	Version      string `json:"version"`
	CPULimit     string `json:"cpu_limit"`
	MemoryLimit  string `json:"memory_limit"`
	CacheVersion int    `json:"cache_version"`
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

	// Runner related data to help marvin push Transformer commit.
	PushBranchName    string `json:"commit_branch"`
	CommitAuthor      string `json:"commit_author"`
	CommitAuthorEmail string `json:"commit_author_email"`
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
	RunID               string             `json:"run_id"`
	RunSerial           string             `json:"run_serial"`
	Config              DSConfig           `json:"config"`
	VCSMeta             TransformerVCSMeta `json:"vcs_meta"`
	DSConfigUpdated     bool               `json:"ds_config_updated"`
	IsTriggeredByRunner bool               `json:"is_triggered_by_runner"`
	Transformer         TransformerInfo    `json:"transformer"`
	Meta                map[string]string  `json:"_meta"`
}

type SSHMeta struct {
	User      string `json:"user"`
	Port      string `json:"port"`
	RemoteURL string `json:"remote_url"`
}

type SSHVerifyRun struct {
	RunID string  `json:"run_id"`
	Keys  Keys    `json:"keys"`
	Meta  SSHMeta `json:"ssh_meta"`
}

type SSHVerifyResult struct {
	RunID  string `json:"run_id"`
	Status int    `json:"status"`
}

type SSHVerifyResultCeleryTask struct {
	ID      string          `json:"id"`
	Task    string          `json:"task"`
	KWArgs  SSHVerifyResult `json:"kwargs"`
	Retries int             `json:"retries"`
}

// CancelCheckRun type is the expected structure of a check cancellation
// task to be recieved
//
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

// PatcherRun type is the contract of a patching job that is used
// by the runner to apply and commit the patches of Autofix.
type PatcherRun struct {
	RunID     string         `json:"run_id"`
	RunSerial string         `json:"run_serial"`
	Keys      Keys           `json:"keys"`
	VCSMeta   PatcherVCSMeta `json:"patcher_vcs_meta"`
	PatchMeta PatchMeta      `json:"patch_meta"`
}

type PatchMeta struct {
	Patches     []PatchData `json:"patches"`
	PatchCommit PatchCommit `json:"patch_commit"`
}

type PatcherVCSMeta struct {
	Repository      string `json:"repository"`
	RemoteURL       string `json:"remote_url"`
	BaseBranch      string `json:"base_branch"`
	BaseOID         string `json:"base_oid"`
	CheckoutOID     string `json:"checkout_oid"`
	CloneSubmodules bool   `json:"clone_submodules"`
}

type PatchData struct {
	Filename  string   `json:"filename"`
	ChangeIDs []string `json:"change_ids"`
	Action    string   `json:"action"`
}

type PatchCommit struct {
	Commit Commit `json:"commit"`
	Author Author `json:"author"`
	Remote Remote `json:"remote"`
}

type Commit struct {
	Title             string `json:"title"`
	Message           string `json:"message"`
	DestinationBranch string `json:"destination_branch"`
	CommitSHA         string `json:"commit_sha"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Remote struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Beacon type is the expected structure of a beacon task
// to be received
//
//proteus:generate
type BeaconRun struct {
	RunID        string `json:"run_id"`
	RepositoryID int64  `json:"repository_id"`
}
