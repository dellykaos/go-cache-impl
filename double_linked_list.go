package cache

import "fmt"

// node store key and value of the cache, and also works as double linked list node
// for eviction policy Least Recently Used or Least Frequently Used
type node struct {
	key   string
	value string
	next  *node
	prev  *node
	freq  int
}

func newNode(key, value string) *node {
	return &node{
		key:   key,
		value: value,
		freq:  1,
	}
}

// doublyLinkedList is data structure that act as double linked list
type doublyLinkedList struct {
	head *node
	tail *node
}

// newDoublyLinkedList initiate new doublyLinkedList data
func newDoublyLinkedList() *doublyLinkedList {
	d := &doublyLinkedList{head: &node{}, tail: &node{}}
	d.head.next = d.tail
	d.tail.prev = d.head

	return d
}

func (d *doublyLinkedList) moveToFront(node *node) {
	d.remove(node)
	d.add(node)
}

func (d *doublyLinkedList) add(node *node) {
	node.next = d.head.next
	node.prev = d.head
	d.head.next.prev = node
	d.head.next = node
}

func (d *doublyLinkedList) remove(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (d *doublyLinkedList) popTail() *node {
	if d.isEmpty() {
		return nil
	}
	lastnode := d.tail.prev
	d.remove(lastnode)
	return lastnode
}

func (d *doublyLinkedList) isEmpty() bool {
	return d.head.next == d.tail
}

func (d *doublyLinkedList) printNode() {
	i := 0
	node := d.head
	msg := "#######\n"
	for node != nil {
		m := node.key
		if node == d.head {
			m = "head"
		}
		if node == d.tail {
			m = "tail"
		}
		msg += fmt.Sprintf("node :%d , key: %s\n", i, m)
		node = node.next
		i++
	}
	msg += "######"
	fmt.Println(msg)
}
