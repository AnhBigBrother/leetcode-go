package main

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	row := make([]bool, n)
	col := make([]bool, n)
	dia1 := map[int]bool{} // row_i - col_i
	dia2 := map[int]bool{} // row_i + col_i
	board := [][]byte{}
	for i := 0; i < n; i++ {
		row := []byte{}
		for j := 0; j < n; j++ {
			row = append(row, '.')
		}
		board = append(board, row)
	}
	var backtrack func(count_q, start_row int, board *[][]byte)
	backtrack = func(count_q, start_row int, board *[][]byte) {
		if count_q == n {
			foo := []string{}
			for _, s := range *board {
				foo = append(foo, string(s))
			}
			ans = append(ans, foo)
			return
		}
		if start_row >= n*n {
			return
		}
		i := start_row
		for j := 0; j < n; j++ {
			if !row[i] && !col[j] && !dia1[i-j] && !dia2[i+j] {
				row[i] = true
				col[j] = true
				dia1[i-j] = true
				dia2[i+j] = true
				(*board)[i][j] = 'Q'
				backtrack(count_q+1, start_row+1, board)
				(*board)[i][j] = '.'
				row[i] = false
				col[j] = false
				dia1[i-j] = false
				dia2[i+j] = false
			}
		}
	}
	backtrack(0, 0, &board)
	return ans
}
