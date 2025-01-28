package main

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	if n == 1 {
		return 0
	}
	left, right := make([]int, n), make([]int, n)
	left[0] = grid[1][0]
	right[n-1] = grid[0][n-1]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + grid[1][i]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + grid[0][i]
	}
	ans := min(left[n-2], right[1])
	for i := 1; i < n-1; i++ {
		ans = min(ans, max(left[i-1], right[i+1]))
	}
	return int64(ans)
}
