package main

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	height := make([][]int, m)
	for i := range height {
		height[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			height[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[i][j] = height[i-1][j] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = max(ans, largestRectangleArea(height[i]))
	}
	return ans
}
