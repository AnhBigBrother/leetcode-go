package main

func maxScore1422(s string) int {
	n := len(s)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = left[i-1]
		if s[i-1] == '0' {
			left[i]++
		}
	}
	if s[n-1] == '1' {
		right[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1]
		if s[i] == '1' {
			right[i]++
		}
	}
	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, left[i]+right[i])
	}
	return ans
}
