package analyzer

type Location struct {
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

type Issue struct {
	IssueCode string   `json:"issue_code"`
	Location  Location `json:"location"`
}

type AnalysisOut struct {
	AnalyzerShortcode string  `json:"analyzer_shortcode"`
	Issues            []Issue `json:"issues"`
}
