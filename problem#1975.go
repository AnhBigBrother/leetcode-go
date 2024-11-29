package main

func maxMatrixSum(matrix [][]int) int64 {
	sum := 0
	count_nev := 0
	mini := 100001
	for _, row := range matrix {
		for _, x := range row {
			if x < 0 {
				count_nev++
				x = -x
			}
			sum += x
			mini = min(mini, x)
		}
	}
	if count_nev%2 == 0 {
		return int64(sum)
	}
	return int64(sum - 2*mini)
}
