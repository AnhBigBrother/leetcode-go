package main

import (
	"slices"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, x := range nums {
		strs[i] = strconv.Itoa(x)
	}
	sortCallback := func(a string, b string) int {
		c := a + b
		d := b + a
		if c < d {
			return 1
		}
		if c > d {
			return -1
		}
		return 0
	}
	slices.SortStableFunc(strs, sortCallback)
	ans := strings.Join(strs, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
