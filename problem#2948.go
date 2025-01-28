package main

import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	idx := []int{}
	for i := 0; i < n; i++ {
		idx = append(idx, i)
	}
	slices.SortFunc(idx, func(i, j int) int {
		if nums[i] > nums[j] {
			return 1
		}
		return -1
	})
	window := []int{idx[0]}
	for i := 1; i < n; i++ {
		if nums[idx[i]]-nums[window[len(window)-1]] <= limit {
			window = append(window, idx[i])
		} else {
			vals := []int{}
			for _, j := range window {
				vals = append(vals, nums[j])
			}
			slices.Sort(window)
			for j := range window {
				nums[window[j]] = vals[j]
			}
			window = []int{idx[i]}
		}
	}
	if len(window) > 1 {
		vals := []int{}
		for _, j := range window {
			vals = append(vals, nums[j])
		}
		slices.Sort(window)
		for j := range window {
			nums[window[j]] = vals[j]
		}
	}
	return nums
}
