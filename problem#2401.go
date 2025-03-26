package main

func longestNiceSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	l, r := 0, 1
	x := nums[0]
	ans := 1
	for r < n {
		for nums[r]&x != 0 && l < r {
			x ^= nums[l]
			l++
		}
		x ^= nums[r]
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
