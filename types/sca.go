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
	SCATarget SCATarget `json:"sca_target"`
	CheckSeq  int	    `json:"check_seq"`
}

type SCARun struct {
	RunID      string             `json:"run_id"`
	RunSerial  int                `json:"run_serial"`
	IsFullRun  bool               `json:"is_full_run"`
	VCSMeta    AnalysisRunVCSMeta `json:"vcs_meta"`
	SCAChecks  []SCACheck         `json:"sca_checks"`
	Keys       Keys               `json:"keys"`
	Meta       map[string]string  `json:"_meta"`
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
}

type Dependency struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	PackageGroup PackageGroup `json:"package_group"`
	PackageType  PackageType  `json:"package_type"`
	Ecosystem    string       `json:"ecosystem"`
	Purl         string       `json:"purl"`
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

	FixedVersions      []string         `json:"fixed_versions"`
	IntroducedVersions []string         `json:"introduced_versions"`
	IntroducedThrough  [][]string       `json:"introduced_through"`
	ReferenceStack     []ReferenceStack `json:"reference_stack"`
}

type Reachability string

const (
	REACHABLE            Reachability = "rea"
	UNREACHABLE          Reachability = "unr"
	UNKNOWN_REACHABILITY Reachability = "unk"
)

type ReferenceStack struct {
	Filepath string   `json:"filepath"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type SCAResult struct {
	RunID     string            `json:"run_id"`
	RunSerial int               `json:"run_serial"`
	CheckSeq  int               `json:"check_seq"`
	Status    Status            `json:"status"`
	Targets   []SCATargetResult `json:"targets"`
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
	RunID     string                 `json:"run_id"`
	RunSerial int                    `json:"run_serial"`
	CheckSeq  int                    `json:"check_seq"`
	Status    Status                 `json:"status"`
	Targets   []SCATargetRemediation `json:"targets"`
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
	VulnerabilityID string    `json:"vulnerability_id"`
	FixPaths        []FixPath `json:"fix_paths"`
}

type FixPath struct {
	Updates []SCAPackageUpdate `json:"updates"`
}

type SCAPackageUpdate struct {
	Package string `json:"package"`
	From    string `json:"from"`
	To      string `json:"to"`
	Risk    Risk   `json:"risk"`
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
	Vulnerability Vulnerability      `json:"vulnerability"`
	Updates       []SCAPackageUpdate `json:"updates"`
}

type SCAPatchResult struct {
	RunID     string `json:"run_id"`
	RunSerial int    `json:"run_serial"`
	Status    Status `json:"status"`
}

type PatchResultCeleryTask struct {
	ID      string         `json:"id"`
	Task    string         `json:"task"`
	KWArgs  SCAPatchResult `json:"kwargs"`
	Retries int            `json:"retries"`
}
