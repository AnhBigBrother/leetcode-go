package main

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	freq := map[int]int{}
	count_distinct := 0
	res := int64(0)
	sum := int64(0)
	l, r := 0, 0
	for r < n {
		if freq[nums[r]] == 0 {
			count_distinct++
		}
		freq[nums[r]]++
		sum += int64(nums[r])
		if r-l+1 == k {
			if count_distinct == k {
				res = max(res, sum)
			}
			freq[nums[l]]--
			if freq[nums[l]] == 0 {
				count_distinct--
			}
			sum -= int64(nums[l])
			l++
		}
		r++
	}
	return res
}
