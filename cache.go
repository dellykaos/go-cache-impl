package cache

import (
	"fmt"
	"math"
)

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

// Get fetch the cache by key from double linked list, if it's not in cache then
// return empty string value
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

type LFUCache struct {
	minFreq  int
	capacity int
	freqMap  map[int]*doublyLinkedList
	cache    map[string]*node
}

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

func (c *LFUCache) Get(key string) string {
	node, found := c.cache[key]
	if !found {
		return ""
	}

	c.updateFrequency(node)
	return node.value
}

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

func (c *LFUCache) Print() {
	fmt.Println("##########\nmin freq: ", c.minFreq)
	for k, m := range c.freqMap {
		fmt.Println("########\nfreq: ", k)
		m.printNode()
	}
	fmt.Println("########DONE#######")
}