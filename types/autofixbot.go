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
	ID           string `json:"id"`
	Type         string `json:"type"`
	RepositoryID string `json:"repository_id"`
}

type BootstrapResult struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	RepositoryID string `json:"repository_id"`
	Status       string `json:"status"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type SessionStartPayload struct {
	ID                 string             `json:"id"`
	Type               string             `json:"type"`
	SessionStartConfig SessionStartConfig `json:"config,omitempty"`
	BootstrapResult    BootstrapResult    `json:"bootstrap_result,omitempty"`
	EnkiGRPCConfig     EnkiGRPCConfig     `json:"enki_grpc_config,omitempty"`
}

type AutofixBotAnalysisConfig struct {
	Languages []string `json:"languages"`
}

type AutofixBotAnalysis struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	RepositoryID string `json:"repository_id"`
	Diff         struct {
		AnalysisID     string
		SourceType     string
		SourceMetadata json.RawMessage
	}
	Config AutofixBotAnalysisConfig `json:"config"`
}
