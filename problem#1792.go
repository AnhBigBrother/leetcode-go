package main

import "container/heap"

type pqItem1792 struct {
	pass  float64
	total float64
	index int
}

type pQueue1792 []*pqItem1792

func (pq pQueue1792) Len() int { return len(pq) }

func (pq pQueue1792) Less(i, j int) bool {
	xi := (pq[i].pass+1)/(pq[i].total+1) - pq[i].pass/pq[i].total
	xj := (pq[j].pass+1)/(pq[j].total+1) - pq[j].pass/pq[j].total
	return xi > xj
}

func (pq pQueue1792) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pQueue1792) Push(x any) {
	n := len(*pq)
	item := x.(*pqItem1792)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pQueue1792) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := pQueue1792{}
	for _, c := range classes {
		pq = append(pq, &pqItem1792{
			pass:  float64(c[0]),
			total: float64(c[1]),
		})
	}
	heap.Init(&pq)
	for extraStudents > 0 {
		x := heap.Pop(&pq).(*pqItem1792)
		ry := (pq[0].pass+1)/(pq[0].total+1) - pq[0].pass/pq[0].total
		for extraStudents > 0 && (x.pass+1)/(x.total+1)-x.pass/x.total >= ry {
			x.pass++
			x.total++
			extraStudents--
		}
		heap.Push(&pq, &pqItem1792{
			pass:  x.pass,
			total: x.total,
		})
	}
	ans := float64(0)
	for _, item := range pq {
		ans += item.pass / item.total
	}
	return ans / float64(len(classes))
}
