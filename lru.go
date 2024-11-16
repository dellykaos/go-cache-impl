package cache

import "math"

// LRUCache is a cache that have Least Recently Used eviction policy
type LRUCache struct {
	cache    map[string]*node
	nodes    *doublyLinkedList
	capacity int
}

// NewLRUCache initiate LRU Cache with defined capacity. If capacity set to 0,
// then it have maximum integer capacity.
func NewLRUCache(capacity int) *LRUCache {
	if capacity == 0 {
		capacity = math.MaxInt
	}
	return &LRUCache{
		cache:    make(map[string]*node),
		capacity: capacity,
		nodes:    newDoublyLinkedList(),
	}
}

// Print print all node of the LRU Cache to get visualization of current
// double linked list data order.
func (c *LRUCache) Print() {
	c.nodes.printNode()
}

// Get fetch the cache by key from double linked list and move the cache
// node to the head, if it's not in cache then return empty string value
func (c *LRUCache) Get(key string) string {
	if node, found := c.cache[key]; found {
		c.nodes.moveToFront(node)
		return node.value
	}

	return ""
}

// Put set the cache with key and value provided, if the cache is already exists
// then it will replace the value of the cache and move the node to the front,
// if not found and the cache is out of capacity, then it will remove last node
// from double linked list and add new node to the head of the double linked list.
func (c *LRUCache) Put(key, value string) {
	if node, found := c.cache[key]; found {
		node.value = value
		c.nodes.moveToFront(node)
		return
	}

	if len(c.cache) >= c.capacity {
		lastNode := c.nodes.tail.prev
		delete(c.cache, lastNode.key)
		c.nodes.remove(lastNode)
	}

	newNode := newNode(key, value)
	c.nodes.add(newNode)
	c.cache[key] = newNode
}
