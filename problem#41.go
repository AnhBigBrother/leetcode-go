package main

func firstMissingPositive(nums []int) int {
	swap := func(i, j int) {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
	n := len(nums)
	for i := range nums {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			swap(i, nums[i]-1)
		}
	}
	for x := 1; x <= n; x++ {
		if nums[x-1] != x {
			return x
		}
	}
	return n + 1
}
