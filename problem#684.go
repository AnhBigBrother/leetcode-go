package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges) + 1
	ds := NewDisjointSet(n)
	ans := []int{-1, -1}
	for _, e := range edges {
		if ds.Check(e[0], e[1]) {
			ans[0] = e[0]
			ans[1] = e[1]
		}
		ds.Union(e[0], e[1])
	}
	return ans
}
