package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	next := make([][]int, n)
	for _, e := range edges {
		if next[e[0]] == nil {
			next[e[0]] = []int{}
		}
		if next[e[1]] == nil {
			next[e[1]] = []int{}
		}
		next[e[0]] = append(next[e[0]], e[1])
		next[e[1]] = append(next[e[1]], e[0])
	}
	sum := []int{}
	for i := 0; i < n; i++ {
		sum = append(sum, -1)
	}
	var calSum func(root int)
	calSum = func(root int) {
		sum[root] = values[root]
		for _, x := range next[root] {
			if sum[x] == -1 {
				calSum(x)
				sum[root] += sum[x]
			}
		}
	}
	calSum(0)
	ans := 0
	visited := make([]bool, n)
	var solve func(root int)
	solve = func(root int) {
		visited[root] = true
		if sum[root]%k == 0 {
			ans++
		}
		for _, x := range next[root] {
			if !visited[x] {
				solve(x)
			}
		}
	}
	solve(0)
	return ans
}
