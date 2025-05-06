package cache

import "container/heap"

type LFUItem struct {
	Key  int
	Val  int
	Freq int
	Last int
	Idx  int
}

type LFUQueue []*LFUItem

func (pq LFUQueue) Len() int { return len(pq) }

func (pq LFUQueue) Less(i, j int) bool {
	if pq[i].Freq == pq[j].Freq {
		return pq[i].Last < pq[j].Last
	}
	return pq[i].Freq < pq[j].Freq
}

func (pq LFUQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Idx = i
	pq[j].Idx = j
}

func (pq *LFUQueue) Push(x any) {
	n := len(*pq)
	item := x.(*LFUItem)
	item.Idx = n
	*pq = append(*pq, item)
}

func (pq *LFUQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Idx = -1
	*pq = old[0 : n-1]
	return item
}

type LFUCache struct {
	Cache LFUQueue
	LMap  map[int]*LFUItem
	Cap   int
	Last  int
}

// Constructor
func NewLFUCache(capacity int) LFUCache {
	x := LFUCache{
		Cache: LFUQueue{},
		LMap:  map[int]*LFUItem{},
		Cap:   capacity,
		Last:  0,
	}
	heap.Init(&(x.Cache))
	return x
}

func (lfu *LFUCache) Get(key int) int {
	x := lfu.LMap[key]
	if x == nil {
		return -1
	}
	lfu.Last++
	x.Last = lfu.Last
	x.Freq++
	heap.Fix(&(lfu.Cache), x.Idx)
	return x.Val
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.LMap[key] == nil {
		if lfu.Cache.Len() == lfu.Cap {
			x := heap.Pop(&(lfu.Cache)).(*LFUItem)
			lfu.LMap[x.Key] = nil
		}
		lfu.Last++
		newItem := LFUItem{
			Key:  key,
			Val:  value,
			Freq: 1,
			Last: lfu.Last,
			Idx:  len(lfu.Cache),
		}
		lfu.LMap[key] = &newItem
		heap.Push(&(lfu.Cache), &newItem)
	} else {
		x := lfu.LMap[key]
		x.Freq++
		lfu.Last++
		x.Last = lfu.Last
		x.Val = value
		heap.Fix(&(lfu.Cache), x.Idx)
	}
}
