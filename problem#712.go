package main

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	minDel := [][]int{} // minDel[i][j] = minimumDeleteSum(s1[:i], s2[:j])
	for range m + 1 {
		minDel = append(minDel, make([]int, n+1))
	}
	for i := 1; i <= m; i++ {
		minDel[i][0] = minDel[i-1][0] + int(s1[i-1])
	}
	for j := 1; j <= n; j++ {
		minDel[0][j] = minDel[0][j-1] + int(s2[j-1])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := int(s1[i-1]) + int(s2[j-1])
			if s1[i-1] == s2[j-1] {
				cost = 0
			}
			minDel[i][j] = min(minDel[i-1][j]+int(s1[i-1]), minDel[i][j-1]+int(s2[j-1]), minDel[i-1][j-1]+cost)
		}
	}
	return minDel[m][n]
}
