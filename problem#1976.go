package main

import (
	"container/heap"
	"math"
)

type Item1976 struct {
	Node int
	Cost int
}

type PQueue1976 []*Item1976

func (pq PQueue1976) Len() int { return len(pq) }

func (pq PQueue1976) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PQueue1976) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue1976) Push(x any) {
	cur := x.(*Item1976)
	*pq = append(*pq, cur)
}

func (pq *PQueue1976) Pop() any {
	old := *pq
	n := len(old)
	cur := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return cur
}

func countPaths(n int, roads [][]int) int {
	MOD := int(1e9 + 7)
	count := make([]int, n)
	count[0] = 1
	minCost := []int{}
	for range n {
		minCost = append(minCost, math.MaxInt64)
	}
	minCost[0] = 0
	next := make([][][2]int, n)
	for _, r := range roads {
		u, v, time := r[0], r[1], r[2]
		next[u] = append(next[u], [2]int{v, time})
		next[v] = append(next[v], [2]int{u, time})
	}
	pq := PQueue1976{}
	pq = append(pq, &Item1976{0, 0})
	heap.Init(&pq)
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item1976)
		if cur.Cost > minCost[cur.Node] || cur.Node == n-1 {
			continue
		}
		for _, n := range next[cur.Node] {
			if cur.Cost+n[1] < minCost[n[0]] {
				minCost[n[0]] = cur.Cost + n[1]
				count[n[0]] = count[cur.Node]
				heap.Push(&pq, &Item1976{
					Node: n[0],
					Cost: minCost[n[0]],
				})
			} else if cur.Cost+n[1] == minCost[n[0]] {
				count[n[0]] += count[cur.Node]
				count[n[0]] %= MOD
			}
		}
	}
	return count[n-1]
}
