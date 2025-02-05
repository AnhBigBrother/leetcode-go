package main

func check(nums []int) bool {
	idx := -1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if idx > -1 {
				return false
			}
			idx = i
		}
	}
	if idx == -1 {
		return true
	}
	for i := idx + 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}

	return nums[len(nums)-1] <= nums[0]
}
