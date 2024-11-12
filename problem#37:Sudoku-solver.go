package main

func solveSudoku(board [][]byte) {
	row, col, box := [][]bool{}, [][]bool{}, [][]bool{}
	for i := 0; i < 9; i++ {
		row = append(row, make([]bool, 10))
		col = append(col, make([]bool, 10))
		box = append(box, make([]bool, 10))
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := (i/3)*3 + j/3
			if board[i][j] != '.' {
				row[i][board[i][j]-'0'] = true
				col[j][board[i][j]-'0'] = true
				box[b][board[i][j]-'0'] = true
			}
		}
	}
	flag := false
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx > 80 || flag {
			flag = true
			return
		}
		i, j := idx/9, idx%9
		b := (i/3)*3 + j/3
		if board[i][j] == '.' {
			for x := '1'; x <= '9'; x++ {
				int_x := int(x - '0')
				if !row[i][int_x] && !col[j][int_x] && !box[b][int_x] {
					board[i][j] = byte(x)
					row[i][int_x] = true
					col[j][int_x] = true
					box[b][int_x] = true
					backtrack(idx + 1)
					if flag {
						return
					}
					board[i][j] = '.'
					row[i][int_x] = false
					col[j][int_x] = false
					box[b][int_x] = false
				}
			}
		} else {
			backtrack(idx + 1)
		}
	}
	backtrack(0)
}

/*Test:
board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
solveSudoku(board)
fmt.Println(board)
*/
