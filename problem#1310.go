package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	xarr := make([]int, len(arr))
	xarr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		xarr[i] = xarr[i-1] ^ arr[i]
	}
	fmt.Println(xarr)
	for i, q := range queries {
		if q[0] == 0 {
			ans[i] = xarr[q[1]]
		} else {
			ans[i] = xarr[q[0]-1] ^ xarr[q[1]]
		}
	}
	return ans
}
