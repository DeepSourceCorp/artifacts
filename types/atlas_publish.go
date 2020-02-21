package types

type Status struct {
	Code     int    `json:"code"`
	HMessage string `json:"hmessage"`
	Err      string `json:"err"`
}

type RepoResult struct {
	RunID    string                 `json:"run_id"`
	Status   Status                 `json:"status"`
	Language map[string]interface{} `json:"language_meta"`
	DSConfig DSConfig               `json:"ds_config"`
}

type RepoResultCeleryTask struct {
	ID      string     `json:"id"`
	Task    string     `json:"task"`
	KWArgs  RepoResult `json:"kwargs"`
	Retries int        `json:"retries"`
}

type DiffMeta struct {
	Additions [][]int `json:"additions"`
	Deletions [][]int `json:"deletions"`
}

type Position struct {
	Begin struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"begin"`
	End struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"end"`
}

type Location struct {
	Path     string   `json:"path"`
	Position Position `json:"position"`
}

type Issue struct {
	IssueCode     string   `json:"issue_code"`
	IssueText     string   `json:"issue_text"`
	Location      Location `json:"location"`
	ProcessedData struct {
		SourceCode struct {
			Rendered string `json:"rendered"`
		} `json:"source_code,omitempty"`
	} `json:"processed_data,omitempty"`
}

// Location of an issue
type IssueLocation struct {
	Path     string `json:"path"`
	Position struct {
		Begin struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"begin"`
		End struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"end"`
	} `json:"position"`
}

type Namespace struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type Metric struct {
	MetricCode string      `json:"metric_code"`
	Namespaces []Namespace `json:"namespaces"`
}

type AnalysisError struct {
	HMessage string `json:"hmessage"`
	Level    int    `json:"level"`
}

type AnalysisReport struct {
	Issues   []Issue         `json:"issues"`
	Metrics  []Metric        `json:"metrics,omitempty"`
	IsPassed bool            `json:"is_passed"`
	Errors   []AnalysisError `json:"errors"`
	FileMeta struct {
		IfAll    bool                `json:"if_all"`
		Deleted  []string            `json:"deleted"`
		Renamed  []string            `json:"renamed"`
		Modified []string            `json:"modified"`
		DiffMeta map[string]DiffMeta `json:"diff_meta,omitempty"`
	} `json:"file_meta"`
	ExtraData interface{} `json:"extra_data"`
}

type Change struct {
	BeforeHTML string `json:"before_html"`
	AfterHTML  string `json:"after_html"`
	Changeset  string `json:"changeset"`
}

type Patch struct {
	Filename string   `json:"filename"`
	Changes  []Change `json:"changes"`
}

type AutofixReport struct {
	CodeDir       string   `json:"code_directory"`
	ModifiedFiles []string `json:"modified_files"`
	IssuesFixed   int      `json:"issues_fixed"`
	Metrics       []Metric `json:"metrics,omitempty"`
	Patches       []Patch  `json:"patches"`
	Errors        []struct {
		HMessage string `json:"hmessage"`
		Level    int    `json:"level"`
	} `json:"errors"`
	ExtraData interface{} `json:"extra_data"`
}

type AnalysisResult struct {
	RunID    string         `json:"run_id"`
	Status   Status         `json:"status"`
	CheckSeq string         `json:"check_seq"`
	Report   AnalysisReport `json:"report"`
}

type AutofixResult struct {
	RunID    string        `json:"run_id"`
	Status   Status        `json:"status"`
	CheckSeq string        `json:"check_seq"`
	Report   AutofixReport `json:"report"`
}

type AnalysisResultCeleryTask struct {
	ID      string         `json:"id"`
	Task    string         `json:"task"`
	KWArgs  AnalysisResult `json:"kwargs"`
	Retries int            `json:"retries"`
}

type AutofixResultCeleryTask struct {
	ID      string        `json:"id"`
	Task    string        `json:"task"`
	KWArgs  AutofixResult `json:"kwargs"`
	Retries int           `json:"retries"`
}

type CancelCheckResult struct {
	RunID  string `json:"run_id"`
	Status Status `json:"status"`
}

type CancelCheckResultCeleryTask struct {
	ID      string            `json:"id"`
	Task    string            `json:"task"`
	KWArgs  CancelCheckResult `json:"kwargs"`
	Retries int               `json:"retries"`
}

type BeaconResult struct {
	Files []struct {
		Path int8 `json:"path"`
		Q1   int8 `json:"q1"`
		Q3   int8 `json:"q2"`
		Q7   int8 `json:"q7"`
		Q15  int8 `json:"q15"`
		Q30  int8 `json:"q30"`
		Q60  int8 `json:"q60"`
		Q180 int8 `json:"a180"`
	} `json:"files"`
}

type BeaconStatusMsg struct {
	RunID     string       `json:"run_id"`
	ProjectID string       `json:"project_id"`
	Status    Status       `json:"status"`
	Result    BeaconResult `json:"result"`
}

type BeaconResultCeleryTask struct {
	ID      string          `json:"id"`
	Task    string          `json:"task"`
	KWArgs  BeaconStatusMsg `json:"kwargs"`
	Retries int             `json:"retries"`
}
