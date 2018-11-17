package types

type DSConfig struct {
	Version         int      `toml:"version" json:"version"`
	ExcludePatterns []string `toml:"exclude_patterns,omitempty" json:"exclude_patterns,omitempty"`
	TestPatterns    []string `toml:"test_patterns,omitempty" json:"test_patterns,omitempty"`
	Analyzers       []struct {
		Name            string   `toml:"name" json:"name"`
		Enabled         bool     `toml:"enabled" json:"enabled"`
		DependencyPaths []string `toml:"dependency_paths,omitempty" json:"dependency_paths,omitempty"`
	} `toml:"analyzers,omitempty" json:"analyzers,omitempty"`
}
