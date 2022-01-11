package types

type MacroBuildContext struct {
	BuildID           string `json:"build_id"`
	AnalyzerShortcode string `json:"analyzer_shortcode"`
	VCSMeta           struct {
		RemoteURL   string `json:"remote_url"`
		BaseBranch  string `json:"base_branch"`
		CheckoutOID string `json:"checkout_oid"`
		TagName     string `json:"tag_name"`
	} `json:"vcs_meta"`
	Keys Keys `json:"keys"`
}

type MacroBuild struct {
	QueueName  string
	RoutingKey string
}
