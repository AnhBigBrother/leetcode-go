package main

import "container/heap"

type Item3342 struct {
	x, y, t  int
	priority int
}

type PriorityQueue3342 []*Item3342

func (pq PriorityQueue3342) Len() int { return len(pq) }

func (pq PriorityQueue3342) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue3342) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue3342) Push(x any) {
	item3342 := x.(*Item3342)
	*pq = append(*pq, item3342)
}

func (pq *PriorityQueue3342) Pop() any {
	old := *pq
	n := len(old)
	item3342 := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item3342
}

func minTimeToReach3342(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	time := [][]int{}
	for range m {
		t := []int{}
		for range n {
			t = append(t, 1e10)
		}
		time = append(time, t)
	}
	time[0][0] = 0
	pq := PriorityQueue3342{}
	heap.Init(&pq)
	heap.Push(&pq, &Item3342{x: 0, y: 0, t: 1, priority: 0})
	for pq.Len() > 0 {
		room := heap.Pop(&pq).(*Item3342)
		x, y, t := room.x, room.y, room.t
		new_t := 1
		if t == 1 {
			new_t = 2
		}
		if x-1 >= 0 && time[x-1][y] > max(time[x][y], moveTime[x-1][y])+t {
			time[x-1][y] = max(time[x][y], moveTime[x-1][y]) + t
			heap.Push(&pq, &Item3342{x: x - 1, y: y, t: new_t, priority: time[x-1][y]})
		}
		if x+1 < m && time[x+1][y] > max(time[x][y], moveTime[x+1][y])+t {
			time[x+1][y] = max(time[x][y], moveTime[x+1][y]) + t
			heap.Push(&pq, &Item3342{x: x + 1, y: y, t: new_t, priority: time[x+1][y]})
		}
		if y-1 >= 0 && time[x][y-1] > max(time[x][y], moveTime[x][y-1])+t {
			time[x][y-1] = max(time[x][y], moveTime[x][y-1]) + t
			heap.Push(&pq, &Item3342{x: x, y: y - 1, t: new_t, priority: time[x][y-1]})
		}
		if y+1 < n && time[x][y+1] > max(time[x][y], moveTime[x][y+1])+t {
			time[x][y+1] = max(time[x][y], moveTime[x][y+1]) + t
			heap.Push(&pq, &Item3342{x: x, y: y + 1, t: new_t, priority: time[x][y+1]})
		}
	}
	return time[m-1][n-1]
}
