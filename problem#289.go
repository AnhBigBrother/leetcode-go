package main

import "fmt"

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < m && y >= 0 && y < n && !(x == i && y == j) {
						if board[x][y]%2 == 1 {
							count++
						}
					}
				}
			}
			if count == 3 && board[i][j] == 0 {
				board[i][j] = 2
			} else if board[i][j] == 1 && (count < 2 || count > 3) {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 3 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
	fmt.Println(board)
}
