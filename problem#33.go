package main

func search(nums []int, target int) int {
	asc_search := func(l, r int) int {
		for l < r {
			m := (l + r) / 2
			if nums[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}
		if nums[l] == target {
			return l
		}
		return -1
	}
	var rotated_search func(l, r int) int
	rotated_search = func(l, r int) int {
		for l < r {
			m := (l + r) / 2
			if nums[l] < nums[m] {
				// [l...m] is in ascending order
				if nums[l] <= target && target <= nums[m] {
					return asc_search(l, m)
				} else {
					return rotated_search(m+1, r)
				}
			} else {
				// [m...r] is in ascending order
				if nums[m+1] <= target && target <= nums[r] {
					return asc_search(m+1, r)
				} else {
					return rotated_search(l, m)
				}
			}
		}
		if nums[l] == target {
			return l
		}
		return -1
	}
	return rotated_search(0, len(nums)-1)
}
