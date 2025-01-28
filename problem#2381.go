package main

import (
	"math"
	"slices"
)

func shiftingLetters(s string, shifts [][]int) string {
	slices.SortFunc[[][]int](shifts, func(a, b []int) int {
		if a[0] == b[0] {
			if a[1] < b[1] {
				return -1
			}
			return 1
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	})
	n := len(s)
	arr := make([]int, n+1)
	for _, s := range shifts {
		if s[2] == 0 {
			arr[s[0]]--
			arr[s[1]+1]++
		} else {
			arr[s[0]]++
			arr[s[1]+1]--
		}
	}
	curr := 0
	ans := []byte(s)
	for i := range s {
		curr += arr[i]
		if curr < -25 {
			x := int(math.Abs(float64(curr)))
			curr = 26 - x%26
		} else if curr > 25 {
			curr = curr % 26
		}
		ans[i] += byte(curr)
		if ans[i] < 'a' {
			ans[i] = 'z' - ('a' - ans[i]) + 1
		} else if ans[i] > 'z' {
			ans[i] = 'a' + (ans[i] - 'z') - 1
		}
	}
	return string(ans)
}
