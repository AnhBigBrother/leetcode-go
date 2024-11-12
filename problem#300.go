package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	subSequence := []int{nums[0]}
	for i := 1; i < n; i++ {
		if subSequence[len(subSequence)-1] < nums[i] {
			subSequence = append(subSequence, nums[i])
		} else {
			l, r := 0, len(subSequence)-1
			for l < r {
				m := (l + r) / 2
				if subSequence[m] < nums[i] {
					l = m + 1
				} else {
					r = m
				}
			}
			subSequence[l] = nums[i]
		}
	}
	return len(subSequence)
}
