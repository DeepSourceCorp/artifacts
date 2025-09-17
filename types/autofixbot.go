package types

type SourceMetadataGCS struct {
	ObjectName          string `json:"object_name,omitempty"`
	BucketName          string `json:"bucket_name,omitempty"`
	UseWorkloadIdentity bool   `json:"use_workload_identity,omitempty"`
	CredentialsJSON     string `json:"credentials_json,omitempty"`
}

type BootstrapConfig struct {
	RepositoryID       string             `json:"repository_id"`
	SourceType         string             `json:"source_type"`
	Size               string             `json:"size,omitempty"`
	Recreate           bool               `json:"recreate,omitempty"`
	SourceMetadata     *SourceMetadataGCS `json:"source_metadata"`
	SnapshotServiceURL string             `json:"snapshot_service_url"`
	CodePath           string             `json:"code_path"`
	PVCName            string             `json:"pvc_name"`
	PatchRepository    bool               `json:"patch_repository,omitempty"`
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
	ID              string           `json:"id"`
	Type            string           `json:"type"`
	BootstrapConfig *BootstrapConfig `json:"config,omitempty"`
	BootstrapResult *BootstrapResult `json:"bootstrap_result,omitempty"`
	EnkiGRPCConfig  *EnkiGRPCConfig  `json:"enki_grpc_config,omitempty"`
}
