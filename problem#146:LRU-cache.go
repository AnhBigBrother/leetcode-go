/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.



Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4


Constraints:

1 <= capacity <= 3000
0 <= key <= 10^4
0 <= value <= 10^5
At most 2 * 10^5 calls will be made to get and put.
*/

package main

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

func (this *LRUCache) Get(key int) int {
	if this.cache[key] == nil {
		return -1
	}
	if this.cache[key] != this.tail {
		node := this.cache[key]
		if node.front != nil {
			node.front.back = node.back
		}
		if node.back != nil {
			node.back.front = node.front
		}
		if node == this.head {
			this.head = node.back
			this.head.front = nil
		}
		this.tail.back = node
		node.front = this.tail
		node.back = nil
		this.tail = node
	}
	return this.cache[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache[key] == nil {
		node := &LRU_Node{key: key, val: value}
		this.cache[key] = node
		if this.length < this.max_length {
			if this.tail == nil {
				this.head = node
				this.tail = node
				this.length++
				return
			}
			this.tail.back = node
			node.front = this.tail
			this.tail = node
			this.length++
		} else {
			if this.max_length == 1 {
				this.cache[this.tail.key] = nil
				this.head = node
				this.tail = node
				return
			}
			this.cache[this.head.key] = nil
			this.head = this.head.back
			this.head.front = nil
			this.tail.back = node
			node.front = this.tail
			this.tail = node
		}
		return
	}
	node := this.cache[key]
	node.val = value
	if node == this.tail {
		return
	}
	if node.front != nil {
		node.front.back = node.back
	}
	if node.back != nil {
		node.back.front = node.front
	}
	if node == this.head {
		this.head = node.back
		this.head.front = nil
	}
	this.tail.back = node
	node.front = this.tail
	node.back = nil
	this.tail = node
}
