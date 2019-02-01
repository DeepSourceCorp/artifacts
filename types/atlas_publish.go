package types

type Status struct {
	Code     int    `json:"code"`
	HMessage string `json:"hmessage"`
	Err      string `json:"err"`
}

type RepoResult struct {
	RunID    string                 `json:"run_id"`
	RunType  string                 `json:"run_type"`
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

type AnalysisArtifacts struct {
	Issues []struct {
		IssueCode string `json:"issue_code"`
		IssueText string `json:"issue_text"`
		Location  struct {
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
		} `json:"location"`
		ProcessedData struct {
			SourceCode struct {
				Begin struct {
					Line int `json:"line"`
				} `json:"begin"`
				End struct {
					Line int `json:"line"`
				} `json:"end"`
				Lines []string `json:"lines"`
			} `json:"source_code,omitempty"`
			IsIgnored bool `json:"is_ignored,omitempty"`
		} `json:"processed_data,omitempty"`
	} `json:"issues"`
	Metrics []struct {
		MetricCode string  `json:"metric_code"`
		Value      float64 `json:"value"`
	} `json:"metrics,omitempty"`
	IsPassed bool `json:"is_passed"`
	FileMeta struct {
		Deleted []string `json:"deleted"`
		Renamed []string `json:"renamed"`
	} `json:"file_meta"`
	ExtraData interface{} `json:"extra_data"`
}

type AnalysisResult struct {
	MachineryTaskID string            `json:"machinery_task_id"`
	RunID           string            `json:"run_id"`
	RunType         string            `json:"run_type"`
	Status          Status            `json:"status"`
	CheckSeq        string            `json:"check_seq"`
	Artifacts       AnalysisArtifacts `json:"artifacts"`
}

type AnalysisResultCeleryTask struct {
	ID      string         `json:"id"`
	Task    string         `json:"task"`
	KWArgs  AnalysisResult `json:"kwargs"`
	Retries int            `json:"retries"`
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
	RunType   string       `json:"beacon"`
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
