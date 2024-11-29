package main

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	count := map[string]int{}
	for i := 0; i < m; i++ {
		s1, s2 := []byte{}, []byte{}
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				s1 = append(s1, '1')
				s2 = append(s2, '0')
			} else {
				s1 = append(s1, '0')
				s2 = append(s2, '1')
			}
		}
		count[string(s1)]++
		count[string(s2)]++
	}
	ans := 1
	for _, x := range count {
		ans = max(ans, x)
	}
	return ans
}
