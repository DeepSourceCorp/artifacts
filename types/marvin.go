package types

//proteus:generate
type MarvinCacheURLs struct {
	MetadataDownload string `toml:"metadataDL"`
	MetadataUpload   string `toml:"metadataUL"`
	CacheDownload    string `toml:"cacheDL"`
	CacheUpload      string `toml:"cacheUL"`
}
