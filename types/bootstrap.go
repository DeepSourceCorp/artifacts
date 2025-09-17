package types

import "encoding/json"

type SourceMetadataGCS struct {
	ObjectName          string `json:"object_name,omitempty"`
	BucketName          string `json:"bucket_name,omitempty"`
	UseWorkloadIdentity bool   `json:"use_workload_identity,omitempty"`
	CredentialsJSON     string `json:"credentials_json,omitempty"`
}

type BootstrapConfig struct {
	RepositoryID       string          `json:"repository_id"`
	SourceType         string          `json:"source_type"`
	SourceMetadata     json.RawMessage `json:"source_metadata"`
	SnapshotServiceURL string          `json:"snapshot_service_url"`
	CodePath           string          `json:"code_path"`
	PVCName            string          `json:"pvc_name"`
	PatchRepository    bool            `json:"patch_repository,omitempty"`
}

type EnkiGRPCConfig struct {
	FlowID       string `json:"flow_id"`
	FlowType     string `json:"flow_type"`
	RepositoryID string `json:"repository_id"`
}

type BootstrapResult struct {
	Status         string          `json:"status"`
	Success        bool            `json:"success"`
	ErrorMessage   string          `json:"error_message,omitempty"`
	EnkiGRPCConfig *EnkiGRPCConfig `json:"enki_grpc_config,omitempty"`
}
