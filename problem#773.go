package main

func slidingPuzzle(board [][]int) int {
	str := []byte{}
	checked := map[string]bool{}
	count := 0
	pArr := []int{}
	adjs := [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	byt := []byte{'0', '1', '2', '3', '4', '5'}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			str = append(str, byt[board[i][j]])
			if board[i][j] == 0 {
				pArr = append(pArr, i*3+j)
			}
		}
	}
	strArr := []string{string(str)}
	checked[string(str)] = true
	for len(pArr) > 0 {
		newPArr := []int{}
		newStrArr := []string{}
		for i := range pArr {
			p := pArr[i]
			str := strArr[i]
			if str == "123450" {
				return count
			}
			for _, x := range adjs[p] {
				if (x >= 0 && x <= 2) || (x >= 3 && x <= 5) {
					newStr := []byte(str)
					newStr[p], newStr[x] = newStr[x], newStr[p]
					if !checked[string(newStr)] {
						checked[string(newStr)] = true
						newPArr = append(newPArr, x)
						newStrArr = append(newStrArr, string(newStr))
					}
				}
			}
		}
		count++
		pArr = newPArr
		strArr = newStrArr
	}
	return -1
}
