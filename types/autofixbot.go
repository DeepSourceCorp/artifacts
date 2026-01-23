package types

import "encoding/json"

type SourceMetadataGCS struct {
	ObjectName          string `json:"object_name,omitempty"`
	BucketName          string `json:"bucket_name,omitempty"`
	UseWorkloadIdentity bool   `json:"use_workload_identity,omitempty"`
	CredentialsJSON     string `json:"credentials_json,omitempty"`
}

type SourceMetadataGit struct {
	RepositoryURL string `json:"repository_url,omitempty"`
	Branch        string `json:"branch,omitempty"`
	CommitSHA     string `json:"commit_sha,omitempty"`
}

type SourceMetadataSignedURL struct {
	RepositorySignedURL string `json:"repository_signed_url"`
	DiffSignedURL       string `json:"diff_signed_url"`
}

type SessionStartConfig struct {
	RepositoryID       string          `json:"repository_id"`
	SourceType         string          `json:"source_type"`
	Size               string          `json:"size,omitempty"`
	Recreate           bool            `json:"recreate,omitempty"`
	SourceMetadata     json.RawMessage `json:"source_metadata"`
	SnapshotServiceURL string          `json:"snapshot_service_url,omitempty"`
	CodePath           string          `json:"code_path,omitempty"`
	PVCName            string          `json:"pvc_name,omitempty"`
	PatchRepository    bool            `json:"patch_repository,omitempty"`
}

type EnkiGRPCConfig struct {
	FlowID       string `json:"flow_id"`
	FlowType     string `json:"flow_type"`
	RepositoryID string `json:"repository_id"`
}

type BootstrapResult struct {
	FlowID       string `json:"flow_id"`
	FlowType     string `json:"flow_type"`
	RepositoryID string `json:"repository_id"`
	Status       string `json:"status"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type SessionStartPayload struct {
	FlowID             string             `json:"flow_id"`
	FlowType           string             `json:"flow_type"`
	SessionStartConfig SessionStartConfig `json:"config,omitempty"`
	BootstrapResult    BootstrapResult    `json:"bootstrap_result,omitempty"`
	EnkiGRPCConfig     EnkiGRPCConfig     `json:"enki_grpc_config,omitempty"`
}

type AutofixBotAnalysisConfig struct {
	Analyzers []string `json:"analyzers"`
}

type MarvinAnalyzerData struct {
	ImageName      string `json:"image_name"`
	ImageTag       string `json:"image_tag"`
	Command        string `json:"command"`
	MinCPULimit    int    `json:"cpu_limit"`
	MaxCPULimit    int    `json:"max_cpu_limit"`
	MinMemoryLimit int    `json:"min_memory_limit"`
	MaxMemoryLimit int    `json:"max_memory_limit"`
}

const (
	FlowTypeTUI    = "tui"
	FlowTypeVendor = "vendor"
)

type AutofixBotAnalysis struct {
	// Set by the config gen engine to denote whether the config has been passed through config gen.
	// This ensures that we do not retry config gen if the number of analyzers and targets detected
	// are zero.
	IsGenerated bool `json:"is_generated"`

	FlowID       string `json:"flow_id"`
	FlowType     string `json:"flow_type"`
	RepositoryID string `json:"repository_id"`

	// Detectors are an optional user supplied list of detectors to use for the
	// analysis.  Only the enabled detectors will be triggered for the analysis.
	Detectors []string `json:"detectors"`

	Fixers []string `json:"fixers"`

	// Languages supplied by the user.  We will still auto detect the languages
	// however, we will exclude any languages that are not in the user supplied
	// list.  If not supplied, we will run all the auto detected languages.
	Languages []string `json:"languages"`

	TestPatterns    []string `json:"test_patterns"`
	ExcludePatterns []string `json:"exclude_patterns"`

	Diff struct {
		AnalysisID string `json:"analysis_id"`

		SourceType string `json:"source_type"`

		// TODO(Vishnu): Remove this and replace with concrete types.
		SourceMetadata json.RawMessage `json:"source_metadata"`

		CommitRange struct {
			FromCommitOID string `json:"from_commit_oid"`
			ToCommitOID   string `json:"to_commit_oid"`
		} `json:"commit_range"`
	} `json:"diff"`

	// These are generated after the config generation step.

	// DSConfig is the generated DSConfig for the analysis.
	DSConfig DSConfig `json:"ds_config"`

	// ExpectedAnalyzers are the final list of analyzers that will be run
	// based on the config.
	ExpectedAnalyzers []string `json:"expected_analyzers"`

	// ExpectedTargets are the final list of SCA targets that will be run
	// based on the config.
	ExpectedTargets []SCATarget `json:"expected_targets"`

	// MarvinAnalyzerDataMap is generated with the analyzer container
	// metadata based on the config generation step.
	MarvinAnalyzerDataMap map[string]MarvinAnalyzerData `json:"marvin_analyzer_data_map,omitempty"`
}

type DetachedRunResult struct {
	FlowID   string `json:"flow_id"`
	FlowType string `json:"flow_type"`
	Status   Status `json:"status"`
}

type DetachedRunResultCeleryTask struct {
	ID      string            `json:"id"`
	Task    string            `json:"task"`
	KWArgs  DetachedRunResult `json:"kwargs"`
	Retries int               `json:"retries"`
}

type MarvinAutofixBotAnalysisConfig struct {
	RunID             string   `toml:"runID"`
	AnalyzerShortcode string   `toml:"analyzerShortcode"`
	AnalyzerCommand   string   `toml:"analyzerCommand"`
	AnalyzerType      string   `toml:"analyzerType"`
	BaseOID           string   `toml:"baseOID"`
	CheckoutOID       string   `toml:"checkoutOID"`
	Repository        string   `toml:"repository"`
	IsFullRun         bool     `toml:"isFullRun"`
	Processors        []string `toml:"processors"`
}
