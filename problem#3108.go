package main

func minimumCost(n int, edges [][]int, query [][]int) []int {
	ds := NewDisjointSet(n)
	for _, e := range edges {
		ds.Union(e[0], e[1])
	}
	weight := map[int]int{}
	for _, e := range edges {
		x := ds.Find(e[0])
		w, ok := weight[x]
		if !ok {
			weight[x] = e[2]
		} else {
			weight[x] = w & e[2]
		}
	}

	ans := []int{}
	for _, q := range query {
		x, y := ds.Find(q[0]), ds.Find(q[1])
		if x == y {
			ans = append(ans, weight[x])
		} else {
			ans = append(ans, -1)
		}
	}

	return ans
}
