package cache

type LRU_Node struct {
	key   int
	val   int
	front *LRU_Node
	back  *LRU_Node
}

type LRUCache struct {
	head       *LRU_Node
	tail       *LRU_Node
	cache      map[int]*LRU_Node
	length     int
	max_length int
}

// Constructor
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		length:     0,
		max_length: capacity,
		cache:      map[int]*LRU_Node{},
	}
}

func (lru *LRUCache) Get(key int) int {
	if lru.cache[key] == nil {
		return -1
	}
	if lru.cache[key] != lru.tail {
		node := lru.cache[key]
		if node.front != nil {
			node.front.back = node.back
		}
		if node.back != nil {
			node.back.front = node.front
		}
		if node == lru.head {
			lru.head = node.back
			lru.head.front = nil
		}
		lru.tail.back = node
		node.front = lru.tail
		node.back = nil
		lru.tail = node
	}
	return lru.cache[key].val
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.cache[key] == nil {
		node := &LRU_Node{key: key, val: value}
		lru.cache[key] = node
		if lru.length < lru.max_length {
			if lru.tail == nil {
				lru.head = node
				lru.tail = node
				lru.length++
				return
			}
			lru.tail.back = node
			node.front = lru.tail
			lru.tail = node
			lru.length++
		} else {
			if lru.max_length == 1 {
				lru.cache[lru.tail.key] = nil
				lru.head = node
				lru.tail = node
				return
			}
			lru.cache[lru.head.key] = nil
			lru.head = lru.head.back
			lru.head.front = nil
			lru.tail.back = node
			node.front = lru.tail
			lru.tail = node
		}
		return
	}
	node := lru.cache[key]
	node.val = value
	if node == lru.tail {
		return
	}
	if node.front != nil {
		node.front.back = node.back
	}
	if node.back != nil {
		node.back.front = node.front
	}
	if node == lru.head {
		lru.head = node.back
		lru.head.front = nil
	}
	lru.tail.back = node
	node.front = lru.tail
	node.back = nil
	lru.tail = node
}
