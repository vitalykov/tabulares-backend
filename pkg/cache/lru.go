package cache

import (
	"errors"
	"sync"
)

// TODO: Add Delete method

type LRUCache[K comparable, V any] struct {
	nodes map[K]*ListNode[V]
	data  *List[V]
	mu    sync.Mutex
	cap   int
}

func NewLRUCache[K comparable, V any](capacity int) (*LRUCache[K, V], error) {
	if capacity <= 0 {
		return nil, errors.New("Capacity must be positive")
	}
	cache := LRUCache[K, V]{
		nodes: make(map[K]*ListNode[V]),
		data:  NewList[V](),
		cap:   capacity,
	}
	return &cache, nil
}

func (lru *LRUCache[K, V]) Size() int {
	return lru.data.size
}

func (lru *LRUCache[K, V]) Get(key K) (V, bool) {
	lru.mu.Lock()
	node, ok := lru.nodes[key]
	var val V
	if ok {
		val = node.val
		lru.data.MoveFront(node)
	}
	lru.mu.Unlock()
	return val, ok
}

func (lru *LRUCache[K, V]) Set(key K, val V) {
	lru.mu.Lock()
	node, ok := lru.nodes[key]
	if ok {
		lru.data.MoveFront(node)
		node.val = val
	} else {
		if lru.data.size == lru.cap {
			// val := lru.data.Back().val
			lru.data.PopBack()
			delete(lru.nodes, key)
		}
		lru.data.PushFront(val)
		lru.nodes[key] = lru.data.Front()
	}
	lru.mu.Unlock()
}

func (lru *LRUCache[K, V]) Delete(key K) {
	lru.mu.Lock()
	node := lru.nodes[key]
	delete(lru.nodes, key)
	lru.data.Delete(node)
	lru.mu.Unlock()
}

func (lru *LRUCache[K, V]) Clear() {
	lru.mu.Lock()
	lru.data = NewList[V]()
	lru.nodes = make(map[K]*ListNode[V])
	lru.mu.Unlock()
}
