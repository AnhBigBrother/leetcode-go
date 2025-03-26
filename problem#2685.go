package main

func countCompleteComponents(n int, edges [][]int) int {
	ds := NewDisjointSet(n)
	countNeighbor := make([]int, n)
	countNodeInSet := make([]int, n)
	isValid := make([]bool, n)
	ans := 0
	for _, e := range edges {
		ds.Union(e[0], e[1])
		countNeighbor[e[0]]++
		countNeighbor[e[1]]++
	}
	for i := range n {
		root := ds.Find(i)
		isValid[root] = true
		countNodeInSet[root]++
	}
	for i := range n {
		root := ds.Find(i)
		if countNodeInSet[root]-1 > countNeighbor[i] {
			isValid[root] = false
		}
	}
	for i := range n {
		if isValid[i] {
			ans++
		}
	}

	return ans
}
