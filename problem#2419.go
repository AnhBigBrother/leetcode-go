package main

func longestSubarray(nums []int) int {
	k := nums[0]
	for _, x := range nums {
		if x > k {
			k = x
		}
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			x := 0
			for i < len(nums) && nums[i] == k {
				x++
				i++
			}
			if res < x {
				res = x
			}
		}
	}
	return res
}
