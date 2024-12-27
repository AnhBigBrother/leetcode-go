package main

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	next := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			next[i] = -1
		} else {
			next[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, q := range queries {
		if q[0] > q[1] {
			q[0], q[1] = q[1], q[0]
		}
	}
	ans := []int{}
	cached := map[int]map[int]int{}
	var find func(a, b int) int
	find = func(a, b int) int {
		if cached[a] == nil {
			cached[a] = map[int]int{}
		} else {
			val, ok := cached[a][b]
			if ok {
				return val
			}
		}
		if b == -1 || heights[b] > heights[a] {
			cached[a][b] = b
			return b
		}
		val := find(a, next[b])
		cached[a][b] = val
		return val
	}
	for _, q := range queries {
		a, b := q[0], q[1]
		if a == b {
			ans = append(ans, a)
			continue
		}
		ans = append(ans, find(a, b))
	}
	return ans
}
