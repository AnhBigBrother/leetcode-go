package main

func findMin153(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[l] < nums[m] {
			if nums[l] > nums[r] {
				l = m + 1
			} else {
				return nums[l]
			}
		} else {
			if nums[m+1] > nums[m] {
				r = m
			} else {
				return nums[m+1]
			}
		}
	}
	return nums[l]
}
