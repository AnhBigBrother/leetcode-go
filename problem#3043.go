package main

import "math"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	countDigit := func(x int) int {
		res := 0
		for x > 0 {
			res++
			x = int(math.Floor(float64(x) / 10))
		}
		return res
	}
	prefix := make(map[int]int)
	for _, x := range arr1 {
		c := countDigit(x)
		for x > 0 {
			prefix[x] = c
			c--
			x = int(math.Floor(float64(x) / 10))
		}
	}
	ans := 0
	for _, x := range arr2 {
		for x > 0 {
			if prefix[x] > 0 {
				ans = max(ans, prefix[x])
			}
			x = int(math.Floor(float64(x) / 10))
		}
	}
	return ans
}
