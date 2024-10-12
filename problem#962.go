package main

func maxWidthRamp(nums []int) int {
	n := len(nums)
	ans := 0
	maxToTheRight := make([]int, n)
	for i, mr := n-1, 0; i >= 0; i-- {
		mr = max(mr, nums[i])
		maxToTheRight[i] = mr
	}
	for l, r := 0, 0; r < n; r++ {
		if maxToTheRight[r] >= nums[l] {
			ans = max(ans, r-l)
		} else {
			l++
		}
	}
	return ans
}
