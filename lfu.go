package cache

import (
	"fmt"
	"math"
)

// LFUCache is a cache that have Least Frequently Used eviction policy
type LFUCache struct {
	minFreq  int
	capacity int
	freqMap  map[int]*doublyLinkedList
	cache    map[string]*node
}

// NewLFUCache initiate LFU Cache with defined capacity. If capacity set to 0,
// then it have maximum integer capacity
func NewLFUCache(capacity int) *LFUCache {
	if capacity == 0 {
		capacity = math.MaxInt
	}
	return &LFUCache{
		minFreq:  0,
		capacity: capacity,
		freqMap:  make(map[int]*doublyLinkedList),
		cache:    make(map[string]*node),
	}
}

func (c *LFUCache) updateFrequency(node *node) {
	freq := node.freq
	c.freqMap[freq].remove(node)
	if c.freqMap[freq].isEmpty() {
		delete(c.freqMap, freq)
		if c.minFreq == freq {
			c.minFreq++
		}
	}

	node.freq++
	freq = node.freq
	if _, found := c.freqMap[freq]; !found {
		c.freqMap[freq] = newDoublyLinkedList()
	}
	c.freqMap[freq].add(node)
}

// Get fetch the cache by key from frequent maps double linked list, and
// increase the frequency node accessed and move the node to increased
// frequent maps double linked list. If it's not in the cache, then
// return empty string value
func (c *LFUCache) Get(key string) string {
	node, found := c.cache[key]
	if !found {
		return ""
	}

	c.updateFrequency(node)
	return node.value
}

// Put set the cache with key and value provided, if the cache is already exists
// then it will replace the value of the cache, increase the frequency node
// accessed and move the node to other frequency map. If the key is not
// exists and the cache is out of capacity, then it will remove last node
// from minimum frequency map double linked list, and add new node to the
// double linked list in the least frequency map.
func (c *LFUCache) Put(key, value string) {
	if c.capacity == 0 {
		return
	}

	node, found := c.cache[key]
	if found {
		node.value = value
		c.updateFrequency(node)
		return
	}

	if len(c.cache) >= c.capacity {
		lfuList := c.freqMap[c.minFreq]
		lfuNode := lfuList.popTail()
		delete(c.cache, lfuNode.key)
		if lfuList.isEmpty() {
			delete(c.freqMap, c.minFreq)
		}
	}

	newNode := newNode(key, value)
	c.cache[key] = newNode
	if _, found := c.freqMap[1]; !found {
		c.freqMap[1] = newDoublyLinkedList()
	}
	c.freqMap[1].add(newNode)
	c.minFreq = 1
}

// Print print all frequency map and nodes of the LRU Cache to get visualization
// of current frequency map and double linked list data order in each
// frequency map
func (c *LFUCache) Print() {
	fmt.Println("##########\nmin freq: ", c.minFreq)
	for k, m := range c.freqMap {
		fmt.Println("########\nfreq: ", k)
		m.printNode()
	}
	fmt.Println("########DONE#######")
}
