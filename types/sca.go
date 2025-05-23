package types

/////////////////////
// Asgard -> Atlas //
/////////////////////

type SCATarget struct {
	Lockfile       string `json:"lockfile"`
	Manifest       string `json:"manifest"`
	Ecosystem      string `json:"ecosystem"`
	PackageManager string `json:"package_manager"`
}

type SCACheck struct {
	StreamKey string    `json:"stream_key"`
	SCATarget SCATarget `json:"sca_target"`
	CheckSeq  int       `json:"check_seq"`
}

type SCARun struct {
	RunID     string             `json:"run_id"`
	RunSerial int                `json:"run_serial"`
	IsFullRun bool               `json:"is_full_run"`
	VCSMeta   AnalysisRunVCSMeta `json:"vcs_meta"`
	SCAChecks []SCACheck         `json:"sca_checks"`
	Keys      Keys               `json:"keys"`
	Meta      map[string]string  `json:"_meta"`
}

type SCARemediationRun struct {
	RunID         string               `json:"run_id"`
	RunSerial     int                  `json:"run_serial"`
	CheckSeq      int                  `json:"check_seq"`
	VCSMeta       AnalysisRunVCSMeta   `json:"vcs_meta"`
	Keys          Keys                 `json:"keys"`
	AutofixConfig SCARemediationConfig `json:"autofix_config"`
	Meta          map[string]string    `json:"_meta"`
}

type SCARemediationConfig struct {
	Targets []SCARemediationTarget `json:"targets"`
}

type SCARemediationTarget struct {
	SCATarget       SCATarget                      `json:"sca_target"`
	Vulnerabilities map[string][]VulnerabilityInfo `json:"vulnerabilities"`
}

type VulnerabilityInfo struct {
	Package            string   `json:"package_name"`
	Version            string   `json:"package_version"`
	Ecosystem          string   `json:"ecosystem"`
	FixedVersions      []string `json:"fixed_versions"`
	IntroducedVersions []string `json:"introduced_versions"`
}

/////////////////////
// Atlas -> Marvin //
/////////////////////

type MarvinSCAConfig struct {
	RunID                      string `toml:"runID"`
	RunSerial                  int    `toml:"runSerial"`
	CheckSeq                   int    `toml:"checkSeq"`
	AnalyzerCommand            string `toml:"analyzerCommand"`
	AnalysisTaskName           string `toml:"analysisTaskName"`
	RemediationCommand         string `toml:"remediationCommand"`
	RemediationTaskName        string `toml:"remediationTaskName"`
	BaseOID                    string `toml:"baseOID"`
	CheckoutOID                string `toml:"checkoutOID"`
	Repository                 string `toml:"repository"`
	IsFullRun                  bool   `toml:"isFullRun"`
	IsForDefaultAnalysisBranch bool   `toml:"isForDefaultAnalysisBranch"`
}

type MarvinSCARemediationConfig struct {
	RunID               string `toml:"runID"`
	RunSerial           int    `toml:"runSerial"`
	CheckSeq            int    `toml:"checkSeq"`
	BaseOID             string `toml:"baseOID"`
	CheckoutOID         string `toml:"checkoutOID"`
	RemediationCommand  string `toml:"remediationCommand"`
	RemediationTaskName string `toml:"remediationTaskName"`
	Repository          string `toml:"repository"`
}

//////////////////////////////////
// Marvin -> Analyzer -> Marvin //
//////////////////////////////////

type SCAConfig struct {
	SCATargets []SCATarget `json:"sca_targets"`
	Files      []string    `json:"files"`
}

type SCATargetResult struct {
	SCATarget       SCATarget       `json:"sca_target"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Dependencies    []Dependency    `json:"dependencies"`
	ExtraData       interface{}     `json:"extra_data"`
	Errors          []AnalysisError `json:"errors"`
	SpdxSbom        string          `json:"spdx_sbom"`
}

type Dependency struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	PackageGroup PackageGroup `json:"package_group"`
	PackageType  PackageType  `json:"package_type"`
	Ecosystem    string       `json:"ecosystem"`
	Purl         string       `json:"purl"`
	Licenses     []string     `json:"licenses"`
}

type PackageGroup string

const (
	DEV           PackageGroup = "dev"
	PROD          PackageGroup = "pro"
	TEST          PackageGroup = "tes"
	DOCUMENTATION PackageGroup = "doc"
	UNKNOWN_GROUP PackageGroup = "unk"
)

type PackageType string

const (
	DIRECT       PackageType = "dir"
	TRANSITIVE   PackageType = "tra"
	UNKNOWN_TYPE PackageType = "unk"
)

type Vulnerability struct {
	Id      string   `json:"id"`
	Aliases []string `json:"aliases"`
	Summary string   `json:"summary"`
	Details string   `json:"details"`

	PublishedAt string `json:"published_at"`
	UpdatedAt   string `json:"updated_at"`
	WithdrawnAt string `json:"withdrawn_at"`

	ReferenceURLs []struct {
		URL  string `json:"url"`
		Type string `json:"type"`
	} `json:"reference_urls"`

	VulnerabilitySeverity []struct {
		Type  string `json:"type"`
		Score string `json:"score"`
	} `json:"vulnerability_severity"`
	PackageSeverity []struct {
		Type  string `json:"type"`
		Score string `json:"score"`
	} `json:"package_severity"`

	Package string `json:"package"`
	Version string `json:"version"`
	// TODO: move it to Dependency struct
	Ecosystem string `json:"ecosystem"`
	Purl      string `json:"purl"`

	DatabaseSpecific  map[string]interface{} `json:"database_specific"`
	EcosystemSpecific map[string]interface{} `json:"ecosystem_specific"`

	Reachability Reachability `json:"reachability"`
	Fixability   Fixability   `json:"fixability"`

	FixedVersions      []string   `json:"fixed_versions"`
	IntroducedVersions []string   `json:"introduced_versions"`
	IntroducedThrough  [][]string `json:"introduced_through"`
	CallPaths          []CallPath `json:"call_paths"`
}

type Reachability string

const (
	REACHABLE            Reachability = "rea"
	UNREACHABLE          Reachability = "unr"
	UNKNOWN_REACHABILITY Reachability = "unk"
)

type Fixability string

const (
	UNFIXABLE      Fixability = "unf"
	GENERATING_FIX Fixability = "gen"
	POSSIBLE_FIX   Fixability = "pos"
	MANUAL_FIX     Fixability = "man"
	AUTO_FIX       Fixability = "aut"
)

type CallPath struct {
	CallFrames []CallFrame `json:"call_frames"`
}

type CallFrame struct {
	Filepath string   `json:"filepath"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type SCAResult struct {
	RunID        string            `json:"run_id"`
	RunSerial    int               `json:"run_serial"`
	CheckSeq     int               `json:"check_seq"`
	Status       Status            `json:"status"`
	Targets      []SCATargetResult `json:"targets"`
	TargetObject string            `json:"target_object"`
}

type LogStreamEntryLevel string

const (
	INFO    LogStreamEntryLevel = "info"
	WARNING LogStreamEntryLevel = "warning"
	ERROR   LogStreamEntryLevel = "error"
	DEBUG   LogStreamEntryLevel = "debug"
)

type LogStreamEntry struct {
	GUID      string              `json:"guid"`
	Level     LogStreamEntryLevel `json:"level"`
	Message   string              `json:"message"`
	Timestamp string              `json:"timestamp"`
}

// LogStreamEvent is the messaged published by Marvin to the
// log stream processors.
type LogStreamEvent struct {
	ID      string                 `json:"id"`
	Task    string                 `json:"task"`
	Args    []LogStreamEntry       `json:"args"`
	KWargs  map[string]interface{} `json:"kwargs"`
	Retries int                    `json:"retries"`
}

//////////////////////
// Marvin -> Asgard //
//////////////////////

type SCAResultCeleryTask struct {
	ID      string    `json:"id"`
	Task    string    `json:"task"`
	KWArgs  SCAResult `json:"kwargs"`
	Retries int       `json:"retries"`
}

///////////////////////////////
// Vulnerability remediation //
///////////////////////////////

type SCARemediationResult struct {
	RunID        string                 `json:"run_id"`
	RunSerial    int                    `json:"run_serial"`
	CheckSeq     int                    `json:"check_seq"`
	Status       Status                 `json:"status"`
	Targets      []SCATargetRemediation `json:"targets"`
	TargetObject string                 `json:"target_object"`
}

type SCATargetRemediation struct {
	SCATarget SCATarget               `json:"target"`
	Packages  []SCAPackageRemediation `json:"packages"`
}

type SCAPackageRemediation struct {
	Package         string                     `json:"package"`
	Version         string                     `json:"version"`
	Ecosystem       string                     `json:"ecosystem"`
	Vulnerabilities []VulnerabilityRemediation `json:"vulnerabilities"`
}

type VulnerabilityRemediation struct {
	VulnerabilityID string     `json:"vulnerability_id"`
	Fixability      Fixability `json:"fixability"`
	FixPaths        []FixPath  `json:"fix_paths"`
}

type FixPath struct {
	Updates       []SCAPackageUpdate     `json:"updates"`
	ExtraData     map[string]interface{} `json:"extra_data"`
	IsRecommended bool                   `json:"is_recommended"`
}

type SCAPackageUpdate struct {
	Package      string       `json:"package"`
	Ecosystem    string       `json:"ecosystem"`
	PackageGroup PackageGroup `json:"package_group"`
	PackageType  PackageType  `json:"package_type"`
	From         string       `json:"from"`
	To           string       `json:"to"`
	Risk         Risk         `json:"risk"`
}

type Risk struct {
	VersionRiskScore       VersionRiskScore       `json:"version_risk_score"`
	DependencyRiskScore    DependencyRiskScore    `json:"dependency_risk_score"`
	CompatibilityRiskScore CompatibilityRiskScore `json:"compatibility_risk_score"`
	OverallRiskScore       OverallRiskScore       `json:"overall_risk_score"`
}

type VersionRiskScore int

const (
	UNKNOWN_VERSION_RISK_SCORE VersionRiskScore = iota
	PATCH
	MINOR
	MAJOR
)

type DependencyRiskScore int

const (
	UNKNOWN_DEPENDENCY_RISK_SCORE DependencyRiskScore = iota
	VULNERABLE_DEP
	TRANSITIVE_DEP
	DIRECT_DEP
)

type CompatibilityRiskScore int

const (
	UNKNOWN_COMPATIBILITY_RISK_SCORE CompatibilityRiskScore = iota
	MEETS_CONSTRAINT
	BREAKS_CONSTRAINT
)

type OverallRiskScore int

const (
	UNKNOWN_OVERALL_RISK_SCORE OverallRiskScore = iota
	LOW
	MEDIUM
	HIGH
)

type SCARemediationResultCeleryTask struct {
	ID      string               `json:"id"`
	Task    string               `json:"task"`
	KWArgs  SCARemediationResult `json:"kwargs"`
	Retries int                  `json:"retries"`
}

//////////////
// Patching //
//////////////

type SCAPatchRun struct {
	RunID       string         `json:"run_id"`
	RunSerial   int            `json:"run_serial"`
	VCSMeta     PatcherVCSMeta `json:"vcs_meta"`
	PatchConfig SCAPatchConfig `json:"patch_config"`
	PatchCommit PatchCommit    `json:"patch_commit"`
}

type MarvinSCAPatchConfig struct {
	RunID         string `toml:"runID"`
	RunSerial     int    `toml:"runSerial"`
	BaseOID       string `toml:"baseOID"`
	CheckoutOID   string `toml:"checkoutOID"`
	PatchCommand  string `toml:"patchCommand"`
	PatchTaskName string `toml:"patchTaskName"`
}

type SCAPatchConfig struct {
	Targets []SCAPatchTarget `json:"targets"`
}

type SCAPatchTarget struct {
	SCATarget SCATarget  `json:"sca_target"`
	Patches   []SCAPatch `json:"patches"`
}

type SCAPatch struct {
	Vulnerability Vulnerability          `json:"vulnerability"`
	Updates       []SCAPackageUpdate     `json:"updates"`
	ExtraData     map[string]interface{} `json:"extra_data"`
}

type SCAPatchResult struct {
	RunID         string `json:"run_id"`
	RunSerial     int    `json:"run_serial"`
	HeadCommitSHA string `json:"head_commit_sha"`
	Status        Status `json:"status"`
}

type PatchResultCeleryTask struct {
	ID      string         `json:"id"`
	Task    string         `json:"task"`
	KWArgs  SCAPatchResult `json:"kwargs"`
	Retries int            `json:"retries"`
}
