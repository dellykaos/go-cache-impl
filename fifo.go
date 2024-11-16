package cache

// FIFOCache is a cache that have First In First Out eviction policy
type FIFOCache struct {
	cache    map[string]string
	list     []string
	capacity int
}

// NewFIFOCache initiate new FIFO Cache with defined capacity.
// If capacity set to 0, then it have 100.000 capacity
func NewFIFOCache(capacity int) *FIFOCache {
	if capacity == 0 {
		capacity = 100_000
	}
	return &FIFOCache{
		cache:    make(map[string]string),
		list:     make([]string, capacity),
		capacity: capacity,
	}
}

// Get fetch the cache with key, if the key doesn't exists, it will return
// empty string value
func (c *FIFOCache) Get(key string) string {
	val, found := c.cache[key]
	if !found {
		return ""
	}

	return val
}

// Put set the cache with key and value provided, if the key is already
// exists, then it will replace the value of the cache. If the key is
// not exists and the cache is out of capacity, then it will remove last
// key from the list and the cache, and add new key at the first list.
func (c *FIFOCache) Put(key, value string) {
	if key == "" {
		return
	}

	_, found := c.cache[key]
	if found {
		c.cache[key] = value
		return
	}

	if c.list[c.capacity-1] != "" {
		lastItem := c.list[c.capacity-1]
		delete(c.cache, lastItem)
	}

	c.list = append([]string{key}, c.list[0:c.capacity-1]...)
	c.cache[key] = value
}
