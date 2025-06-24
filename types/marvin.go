package types

//proteus:generate
type MarvinCacheURLs struct {
	Enabled bool `toml:"enabled"`

	MetadataDownload string `toml:"metadataDL"`
	MetadataUpload   string `toml:"metadataUL"`
	CacheDownload    string `toml:"cacheDL"`
	CacheUpload      string `toml:"cacheUL"`
	PlaintextCacheDownload string `toml:"plaintextCacheDL"`
	PlaintextCacheUpload string `toml:"plaintextCacheUL"`
}
