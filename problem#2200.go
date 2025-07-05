package main

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	next_k_idx := make([]int, n)
	k_idx := n + 1
	for i := n - 1; i >= 0; i-- {
		if nums[i] == key {
			k_idx = i
		}
		next_k_idx[i] = k_idx
	}
	ans := []int{}
	for i := range n {
		l, r := max(0, i-k), min(n-1, i+k)
		idx := next_k_idx[l]
		if idx <= r {
			ans = append(ans, i)
		}
	}
	return ans
}
