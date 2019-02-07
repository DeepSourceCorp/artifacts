package types

// DSConfig is the struct for .deepsource.toml file
type DSConfig struct {
	Version         int      `toml:"version" json:"version"`
	ExcludePatterns []string `toml:"exclude_patterns,omitempty" json:"exclude_patterns,omitempty"`
	TestPatterns    []string `toml:"test_patterns,omitempty" json:"test_patterns,omitempty"`
	Analyzers       []struct {
		Name                string      `toml:"name" json:"name"`
		RuntimeVersion      string      `toml:"runtime_version,omitempty" json:"runtime_version,omitempty"`
		Enabled             bool        `toml:"enabled" json:"enabled"`
		DependencyFilePaths []string    `toml:"dependency_file_paths,omitempty" json:"dependency_file_paths,omitempty"`
		Meta                interface{} `toml:"meta,omitempty" json:"meta,omitempty"`
	} `toml:"analyzers,omitempty" json:"analyzers,omitempty"`
}
