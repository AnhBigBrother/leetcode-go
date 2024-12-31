package main

// ----- solution1: dynamic programming (optimized from solution2)----------
func numWays(words []string, target string) int {
	m, n := len(words[0]), len(target)
	if m < n {
		return 0
	}
	mod := int(1e9 + 7)
	count := []map[byte]int{}
	for i := 0; i < m; i++ {
		count = append(count, map[byte]int{})
	}
	for _, w := range words {
		for i := 0; i < m; i++ {
			count[i][w[i]]++
		}
	}
	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
		dp[i][n-1] = count[i][target[n-1]]
	}
	for j := n - 2; j >= 0; j-- {
		x := 0
		for i := m - 2; i >= 0; i-- {
			x = (x + dp[i+1][j+1]) % mod
			dp[i][j] = (x * count[i][target[j]]) % mod
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = (ans + dp[i][0]) % mod
	}
	return ans
}

// ----- solution2: backtracking + caching ----------
// func numWays(words []string, target string) int {
// 	n := len(words[0])
// 	mod := int(1e9 + 7)
// 	if len(target) > n {
// 		return 0
// 	}
// 	count := []map[byte]int{}
// 	for i := 0; i < n; i++ {
// 		count = append(count, map[byte]int{})
// 	}
// 	for _, w := range words {
// 		for i := 0; i < n; i++ {
// 			count[i][w[i]]++
// 		}
// 	}
// 	cached := []map[int]int{}
// 	for i := 0; i < n; i++ {
// 		cached = append(cached, map[int]int{})
// 	}
// 	ans := 0
// 	var backtrack func(i, t int) int
// 	backtrack = func(i, t int) int {
// 		if t == len(target)-1 {
// 			cached[i][t] = count[i][target[t]] % mod
// 			return count[i][target[t]] % mod
// 		}
// 		if n-i < len(target)-t {
// 			return 0
// 		}
// 		if val, ok := cached[i][t]; ok {
// 			return val
// 		}
// 		x := 0
// 		for j := i + 1; j < n; j++ {
// 			x = (x + backtrack(j, t+1)) % mod
// 		}
// 		x = (x * count[i][target[t]]) % mod
// 		cached[i][t] = x
// 		return x
// 	}
// 	for i := 0; i < n; i++ {
// 		ans = (ans + backtrack(i, 0)) % mod
// 	}
// 	return ans
// }
