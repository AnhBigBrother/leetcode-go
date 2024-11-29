package main

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	ans := [][]byte{}
	for i := 0; i < n; i++ {
		row := []byte{}
		for j := 0; j < m; j++ {
			row = append(row, '.')
		}
		ans = append(ans, row)
	}
	for i := 0; i < m; i++ {
		p := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '#' {
				ans[p][m-1-i] = '#'
				p--
			} else if box[i][j] == '*' {
				p = j - 1
				ans[j][m-1-i] = '*'
			}
		}
	}
	return ans
}
