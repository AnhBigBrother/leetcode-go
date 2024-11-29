package main

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	ans := len(nums) + 1
	l, r := 0, 0
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
