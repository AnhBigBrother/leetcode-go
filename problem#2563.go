package main

import (
	"slices"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	var count int64 = 0
	for i, j := 0, len(nums)-1; i < j; i++ {
		for i < j && nums[i]+nums[j] > upper {
			j--
		}
		count += int64(j - i) // Add all pairs with sum <= upper
	}
	for i, j := 0, len(nums)-1; i < j; i++ {
		for i < j && nums[i]+nums[j] >= lower {
			j--
		}
		count -= int64(j - i) // Subtract all pairs with sum < lower
	}
	return count
}

// Other solution using binary search:
/* func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	findLow := func(nums []int, idx int) int {
		l, r := 0, idx-1
		for l < r {
			m := (l + r) / 2
			if nums[m]+nums[idx] < lower {
				l = m + 1
			} else {
				r = m
			}
		}
		if nums[l]+nums[idx] < lower {
			return -1
		}
		return l
	}
	findUp := func(nums []int, idx int) int {
		l, r := 0, idx-1
		for l < r {
			m := (l + r + 1) / 2
			if nums[m]+nums[idx] > upper {
				r = m - 1
			} else {
				l = m
			}
		}
		if nums[r]+nums[idx] > upper {
			return -1
		}
		return r
	}

	var ans int64 = 0

	for i := 1; i < len(nums); i++ {
		l := findLow(nums, i)
		u := findUp(nums, i)
		if l == -1 || u == -1 {
			continue
		}
		ans += int64(u - l + 1)
	}

	return ans
} */
