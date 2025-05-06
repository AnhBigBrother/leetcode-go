package main

func countSubarrays2302(nums []int, k int64) int64 {
	var sum int64 = 0
	l, r := int64(0), int64(0)
	var ans int64 = 0
	for ; r < int64(len(nums)); r++ {
		sum += int64(nums[r])
		for sum*(r-l+1) >= k {
			sum -= int64(nums[l])
			l++
		}
		ans += r - l + 1
	}
	return ans
}
