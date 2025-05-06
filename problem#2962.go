package main

func countSubarrays2962(nums []int, k int) int64 {
	maxNum := nums[0]
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
	}
	count := 0
	ans := int64(0)
	l := int64(0)
	for _, n := range nums {
		if n == maxNum {
			count++
		}
		for count == k {
			if nums[l] == maxNum {
				count--
			}
			l++
		}
		ans += l
	}
	return ans
}
