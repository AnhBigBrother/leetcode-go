package main

import "container/heap"

type item_632 struct {
	list        int
	value       int
	idx_in_list int
	index       int
}

type priority_queue_632 []*item_632

func (pq priority_queue_632) Len() int { return len(pq) }

func (pq priority_queue_632) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq priority_queue_632) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priority_queue_632) Push(x any) {
	n := len(*pq)
	it := x.(*item_632)
	it.index = n
	*pq = append(*pq, it)
}

func (pq *priority_queue_632) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func smallestRange(nums [][]int) []int {
	n := len(nums)
	ans := []int{}
	pq := priority_queue_632{}
	heap.Init(&pq)
	M := nums[0][0]
	for i := 0; i < n; i++ {
		heap.Push(&pq, &item_632{
			value:       nums[i][0],
			list:        i,
			idx_in_list: 0,
		})
		M = max(M, nums[i][0])
	}
	ans = append(ans, pq[0].value, M)
	for {
		x := heap.Pop(&pq).(*item_632)
		list := nums[x.list]
		if ans[1]-ans[0] > M-x.value {
			ans[0], ans[1] = x.value, M
		}
		if x.idx_in_list+1 >= len(list) {
			break
		}
		heap.Push(&pq, &item_632{
			value:       list[x.idx_in_list+1],
			list:        x.list,
			idx_in_list: x.idx_in_list + 1,
		})
		M = max(M, list[x.idx_in_list+1])
	}
	return ans
}

/*
Test:
fmt.Println(smallestRange([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}))
*/
