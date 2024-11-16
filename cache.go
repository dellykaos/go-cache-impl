package cache

// Cache is a cache contract
type Cache interface {
	Get(key string) string
	Put(key, value string)
}
