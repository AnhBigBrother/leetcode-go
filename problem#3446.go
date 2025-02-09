package main

import "slices"

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	getDiagonal := func(i, j int) []*int {
		ans := []*int{}
		for i < n && j < n {
			ans = append(ans, &grid[i][j])
			i++
			j++
		}
		return ans
	}
	for i := 0; i < n; i++ {
		arr := getDiagonal(i, 0)
		arr2 := []int{}
		for _, a := range arr {
			arr2 = append(arr2, *a)
		}
		slices.SortFunc(arr2, func(i, j int) int {
			if i < j {
				return 1
			}
			return -1
		})
		for i := range arr {
			*(arr[i]) = arr2[i]
		}
	}
	for j := 1; j < n; j++ {
		arr := getDiagonal(0, j)
		arr2 := []int{}
		for _, a := range arr {
			arr2 = append(arr2, *a)
		}
		slices.Sort(arr2)
		for i := range arr {
			*(arr[i]) = arr2[i]
		}
	}
	return grid
}
