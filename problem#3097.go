package main

import (
	"math"
	"strconv"
)

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	ans := n + 1
	bits := make([]int, 18)
	push := func(x int) {
		s := strconv.FormatInt(int64(x), 2)
		for i := 1; i <= len(s); i++ {
			if s[len(s)-i] == '1' {
				bits[len(bits)-i]++
			}
		}
	}
	pop := func(x int) {
		s := strconv.FormatInt(int64(x), 2)
		for i := 1; i <= len(s); i++ {
			if s[len(s)-i] == '1' {
				bits[len(bits)-i]--
			}
		}
	}
	getValue := func() int {
		res := float64(0)
		for i := 1; i < len(bits); i++ {
			if bits[len(bits)-i] > 0 {
				res += math.Pow(2, float64(i-1))
			}
		}
		return int(res)
	}
	l := 0
	for r, x := range nums {
		push(x)
		for l <= r && getValue() >= k {
			ans = min(ans, r-l+1)
			pop(nums[l])
			l++
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}
