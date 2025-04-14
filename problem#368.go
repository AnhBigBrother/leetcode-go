package main

import "slices"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	slices.Sort(nums)
	ans := []int{}
	next := make([]int, n)
	s_len := make([]int, n)
	s_len[n-1] = 1
	next[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		s_len[i] = 1
		next[i] = -1
		for j := i + 1; j < n; j++ {
			if nums[j]%nums[i] == 0 && s_len[j]+1 > s_len[i] {
				s_len[i] = s_len[j] + 1
				next[i] = j
			}
		}
	}
	max_len := 0
	for i := range n {
		if max_len < s_len[i] {
			max_len = s_len[i]
		}
	}
	x := 0
	for i := range n {
		if s_len[i] == max_len {
			x = i
			break
		}
	}
	for x != -1 {
		ans = append(ans, nums[x])
		x = next[x]
	}
	return ans
}
