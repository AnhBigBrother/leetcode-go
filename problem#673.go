package main

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	lis := make([]int, n) // length of longest subsequence end at i
	cnt := make([]int, n) // count of longest subsequence of of length lis[i]
	for i := range n {
		lis[i] = 1
		cnt[i] = 1
	}
	max_len := 1
	for i := 1; i < n; i++ {
		for j := range i {
			if nums[i] > nums[j] {
				if lis[j]+1 > lis[i] {
					lis[i] = lis[j] + 1
					cnt[i] = cnt[j]
					continue
				}
				if lis[j]+1 == lis[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		max_len = max(max_len, lis[i])
	}
	ans := 0
	for i, l := range lis {
		if l == max_len {
			ans += cnt[i]
		}
	}
	return ans
}
