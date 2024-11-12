package main

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	swap := make([]int, n)
	swap[0] = 1
	no_swap := make([]int, n)
	for i := 1; i < n; i++ {
		no_swap[i], swap[i] = n, n
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			no_swap[i] = no_swap[i-1]
			swap[i] = swap[i-1] + 1
		}
		if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
			no_swap[i] = min(swap[i-1], no_swap[i])
			swap[i] = min(no_swap[i-1]+1, swap[i])
		}
	}
	return min(swap[n-1], no_swap[n-1])
}
