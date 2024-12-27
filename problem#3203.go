package main

import "math"

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	findDiameter := func(n int, next []map[int]bool) int {
		if n <= 2 {
			return n - 1
		}
		leaves := []int{}
		for i, node := range next {
			if len(node) == 1 {
				leaves = append(leaves, i)
			}
		}
		ans := 0
		branchLen := make([]int, n)
		for i := 0; i < len(leaves); i++ {
			leaf := leaves[i]
			for key := range next[leaf] {
				ans = max(ans, branchLen[key]+branchLen[leaf]+1)
				branchLen[key] = max(branchLen[key], branchLen[leaf]+1)
				delete(next[key], leaf)
				if len(next[key]) == 1 {
					leaves = append(leaves, key)
				}
			}
		}
		return ans
	}
	buildTree := func(n int, ed [][]int) []map[int]bool {
		ans := make([]map[int]bool, n)
		for _, e := range ed {
			if ans[e[0]] == nil {
				ans[e[0]] = map[int]bool{}
			}
			if ans[e[1]] == nil {
				ans[e[1]] = map[int]bool{}
			}
			ans[e[0]][e[1]] = true
			ans[e[1]][e[0]] = true
		}
		return ans
	}
	n1 := len(edges1) + 1
	n2 := len(edges2) + 1
	next1 := buildTree(n1, edges1)
	next2 := buildTree(n2, edges2)
	d1 := findDiameter(n1, next1)
	d2 := findDiameter(n2, next2)
	r1 := int(math.Ceil(float64(d1) / 2))
	r2 := int(math.Ceil(float64(d2) / 2))
	return max(d1, d2, r1+r2+1)
}
