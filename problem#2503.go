package main

import "container/heap"

type PQItem2503 struct {
	Value int
	X     int
	Y     int
}

type PriorityQueue2503 []*PQItem2503

func (pq PriorityQueue2503) Len() int { return len(pq) }

func (pq PriorityQueue2503) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

func (pq PriorityQueue2503) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue2503) Push(x any) {
	item := x.(*PQItem2503)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue2503) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func maxPoints2503(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	ans := []int{}
	q_max := queries[0]
	for _, q := range queries {
		if q > q_max {
			q_max = q
		}
	}
	visited := [][]bool{}
	for range m {
		visited = append(visited, make([]bool, n))
	}
	visited[0][0] = true
	count := make([]int, q_max+1)
	pq := PriorityQueue2503{&PQItem2503{
		X: 0, Y: 0, Value: grid[0][0],
	}}
	heap.Init(&pq)
	q := 1
	for q <= q_max {
		if pq.Len() == 0 || pq[0].Value >= q {
			q++
			if q <= q_max {
				count[q] = count[q-1]
			}
			continue
		}
		item := heap.Pop(&pq).(*PQItem2503)
		count[q]++
		x, y := item.X, item.Y
		if x-1 >= 0 && !visited[x-1][y] {
			visited[x-1][y] = true
			heap.Push(&pq, &PQItem2503{
				X:     x - 1,
				Y:     y,
				Value: grid[x-1][y],
			})
		}
		if x+1 < m && !visited[x+1][y] {
			visited[x+1][y] = true
			heap.Push(&pq, &PQItem2503{
				X:     x + 1,
				Y:     y,
				Value: grid[x+1][y],
			})
		}
		if y-1 >= 0 && !visited[x][y-1] {
			visited[x][y-1] = true
			heap.Push(&pq, &PQItem2503{
				X:     x,
				Y:     y - 1,
				Value: grid[x][y-1],
			})
		}
		if y+1 < n && !visited[x][y+1] {
			visited[x][y+1] = true
			heap.Push(&pq, &PQItem2503{
				X:     x,
				Y:     y + 1,
				Value: grid[x][y+1],
			})
		}
	}
	for _, q := range queries {
		ans = append(ans, count[q])
	}
	return ans
}
