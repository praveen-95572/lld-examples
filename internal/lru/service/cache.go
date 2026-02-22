package service

import (
	"sync"
)

// Node represents a doubly linked list node
type Node struct {
	Key   string
	Value interface{}
	prev  *Node
	next  *Node
}

// LRUCache is the main cache structure
type LRUCache struct {
	capacity int
	cache    map[string]*Node
	head     *Node
	tail     *Node
	mutex    sync.RWMutex
}

// NewLRUCache creates a new LRU cache with given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node),
	}
}

// Get retrieves a value and marks it as recently used
func (l *LRUCache) Get(key string) (interface{}, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if node, ok := l.cache[key]; ok {
		l.moveToHead(node)
		return node.Value, true
	}
	return nil, false
}

// Put inserts or updates a key-value pair
func (l *LRUCache) Put(key string, value interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if node, ok := l.cache[key]; ok {
		node.Value = value
		l.moveToHead(node)
		return
	}

	newNode := &Node{
		Key:   key,
		Value: value,
	}

	l.cache[key] = newNode
	l.addToHead(newNode)

	if len(l.cache) > l.capacity {
		// remove LRU item
		delete(l.cache, l.tail.Key)
		l.removeTail()
	}
}

// Internal helper: move node to head
func (l *LRUCache) moveToHead(node *Node) {
	if l.head == node {
		return
	}
	l.removeNode(node)
	l.addToHead(node)
}

// Internal helper: add node to head
func (l *LRUCache) addToHead(node *Node) {
	node.prev = nil
	node.next = l.head
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

// Internal helper: remove node from list
func (l *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}

// Internal helper: remove tail node
func (l *LRUCache) removeTail() {
	if l.tail != nil {
		l.removeNode(l.tail)
	}
}
