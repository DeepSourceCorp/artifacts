package types

//proteus:generate
type AnalyzerIssue struct {
	Shortcode string `json:"shortcode"`
}

//proteus:generate
type MacrobuildRepoData struct {
	Issues       []AnalyzerIssue `json:"issues"`
	AnalyzerData struct {
		Name string `json:"name"`
	} `json:"analyzer_data"`
}

//proteus:generate
type MacrobuildResult struct {
	BuildID  string             `json:"build_id"`
	Status   Status             `json:"status"`
	RepoData MacrobuildRepoData `json:"repo_data"`
}

//proteus:generate
type MacrobuildResultCeleryTask struct {
	ID      string           `json:"id"`
	Task    string           `json:"task"`
	KWArgs  MacrobuildResult `json:"kwargs"`
	Retries int              `json:"retries"`
}
