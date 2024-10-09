package types

/////////////////////
// Asgard -> Atlas //
/////////////////////

type SCATarget struct {
	Lockfile string `json:"lockfile"`
	Manifest string `json:"manifest"`
}

type SCARun struct {
	RunID      string             `json:"run_id"`
	RunSerial  int                `json:"run_serial"`
	CheckSeq   int                `json:"check_seq"`
	IsFullRun  bool               `json:"is_full_run"`
	VCSMeta    AnalysisRunVCSMeta `json:"vcs_meta"`
	SCATargets []SCATarget        `json:"sca_targets"`
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
	Name      string `json:"name"`
	Version   string `json:"version"`
	IsDirect  bool   `json:"is_direct"`
	IsDev     bool   `json:"is_dev"`
	Ecosystem string `json:"ecosystem"`
	Purl      string `json:"purl"`
}

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
	//TODO: move it to Dependency struct
	Ecosystem string `json:"ecosystem"`
	Purl      string `json:"purl"`

	DatabaseSpecific  map[string]interface{} `json:"database_specific"`
	EcosystemSpecific map[string]interface{} `json:"ecosystem_specific"`

	IsReachable        bool `json:"is_reachable"`
	IsFixAvailable     bool `json:"is_fix_available"`
	IsAutofixAvailable bool `json:"is_autofix_available"`

	FixedVersions     []string         `json:"fixed_versions"`
	IntroducedThrough [][]string       `json:"introduced_through"`
	ReferenceStack    []ReferenceStack `json:"reference_stack"`
}

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

type RemediationResult struct {
	RunID     string               `json:"run_id"`
	RunSerial int                  `json:"run_serial"`
	CheckSeq  int                  `json:"check_seq"`
	Status    Status               `json:"status"`
	Target    SCATarget            `json:"target"`
	Packages  []PackageRemediation `json:"packages"`
}

type PackageRemediation struct {
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
	Updates []PackageUpdate `json:"updates"`
}

type PackageUpdate struct {
	Package string `json:"package"`
	From    string `json:"from"`
	To      string `json:"to"`
	Risk    string `json:"risk"`
}

type RemediationResultCeleryTask struct {
	ID      string            `json:"id"`
	Task    string            `json:"task"`
	KWArgs  RemediationResult `json:"kwargs"`
	Retries int               `json:"retries"`
}
