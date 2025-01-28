package main

import "container/heap"

type Item407 struct {
	x   int
	y   int
	val int
}

type PQ407 []*Item407

func (pq PQ407) Len() int { return len(pq) }

func (pq PQ407) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PQ407) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ407) Push(x any) {
	item := x.(*Item407)
	*pq = append(*pq, item)
}

func (pq *PQ407) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	pq := PQ407{}
	visited := [][]bool{}
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}
	for i := 0; i < n; i++ {
		pq = append(pq, &Item407{
			x:   0,
			y:   i,
			val: heightMap[0][i],
		})
		visited[0][i] = true
	}
	if m > 1 {
		for i := 0; i < n; i++ {
			pq = append(pq, &Item407{
				x:   m - 1,
				y:   i,
				val: heightMap[m-1][i],
			})
			visited[m-1][i] = true
		}
	}
	for i := 1; i < m-1; i++ {
		pq = append(pq, &Item407{
			x:   i,
			y:   0,
			val: heightMap[i][0],
		})
		visited[i][0] = true
	}
	if n > 1 {
		for i := 1; i < m-1; i++ {
			pq = append(pq, &Item407{
				x:   i,
				y:   n - 1,
				val: heightMap[i][n-1],
			})
			visited[i][n-1] = true
		}
	}
	heap.Init(&pq)
	ans := 0
	level := pq[0].val
	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*Item407)
		level = max(level, cell.val)
		if cell.x-1 >= 0 && !visited[cell.x-1][cell.y] {
			if heightMap[cell.x-1][cell.y] < level {
				ans += level - heightMap[cell.x-1][cell.y]
			}
			visited[cell.x-1][cell.y] = true
			heap.Push(&pq, &Item407{
				x:   cell.x - 1,
				y:   cell.y,
				val: heightMap[cell.x-1][cell.y],
			})
		}
		if cell.x+1 < m && !visited[cell.x+1][cell.y] {
			if heightMap[cell.x+1][cell.y] < level {
				ans += level - heightMap[cell.x+1][cell.y]
			}
			visited[cell.x+1][cell.y] = true
			heap.Push(&pq, &Item407{
				x:   cell.x + 1,
				y:   cell.y,
				val: heightMap[cell.x+1][cell.y],
			})
		}
		if cell.y-1 >= 0 && !visited[cell.x][cell.y-1] {
			if heightMap[cell.x][cell.y-1] < level {
				ans += level - heightMap[cell.x][cell.y-1]
			}
			visited[cell.x][cell.y-1] = true
			heap.Push(&pq, &Item407{
				x:   cell.x,
				y:   cell.y - 1,
				val: heightMap[cell.x][cell.y-1],
			})
		}
		if cell.y+1 < n && !visited[cell.x][cell.y+1] {
			if heightMap[cell.x][cell.y+1] < level {
				ans += level - heightMap[cell.x][cell.y+1]
			}
			visited[cell.x][cell.y+1] = true
			heap.Push(&pq, &Item407{
				x:   cell.x,
				y:   cell.y + 1,
				val: heightMap[cell.x][cell.y+1],
			})
		}
	}
	return ans
}
