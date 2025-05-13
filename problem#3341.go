package main

import "container/heap"

type Item3341 struct {
	x, y     int
	priority int
}

type PriorityQueue3341 []*Item3341

func (pq PriorityQueue3341) Len() int { return len(pq) }

func (pq PriorityQueue3341) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue3341) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue3341) Push(x any) {
	item3341 := x.(*Item3341)
	*pq = append(*pq, item3341)
}

func (pq *PriorityQueue3341) Pop() any {
	old := *pq
	n := len(old)
	item3341 := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item3341
}

func minTimeToReach3341(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	time := [][]int{}
	for range m {
		row := []int{}
		for range n {
			row = append(row, 1e10)
		}
		time = append(time, row)
	}
	time[0][0] = 0
	pq := PriorityQueue3341{}
	heap.Init(&pq)
	heap.Push(&pq, &Item3341{
		x: 0, y: 0, priority: 0,
	})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item3341)
		x, y := item.x, item.y
		if x-1 >= 0 && time[x-1][y] > max(time[x][y], moveTime[x-1][y])+1 {
			time[x-1][y] = max(time[x][y], moveTime[x-1][y]) + 1
			heap.Push(&pq, &Item3341{x - 1, y, time[x-1][y]})
		}
		if x+1 < m && time[x+1][y] > max(time[x][y], moveTime[x+1][y])+1 {
			time[x+1][y] = max(time[x][y], moveTime[x+1][y]) + 1
			heap.Push(&pq, &Item3341{x + 1, y, time[x+1][y]})
		}
		if y-1 >= 0 && time[x][y-1] > max(time[x][y], moveTime[x][y-1])+1 {
			time[x][y-1] = max(time[x][y], moveTime[x][y-1]) + 1
			heap.Push(&pq, &Item3341{x, y - 1, time[x][y-1]})
		}
		if y+1 < n && time[x][y+1] > max(time[x][y], moveTime[x][y+1])+1 {
			time[x][y+1] = max(time[x][y], moveTime[x][y+1]) + 1
			heap.Push(&pq, &Item3341{x, y + 1, time[x][y+1]})
		}
	}
	return time[m-1][n-1]
}
