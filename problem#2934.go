package main

func minOperations2934(nums1 []int, nums2 []int) int {
	n := len(nums1)
	nums1X := []int{}
	nums2X := []int{}
	nums1X = append(nums1X, nums1...)
	nums2X = append(nums2X, nums2...)
	nums1X[n-1] = nums2[n-1]
	nums2X[n-1] = nums1[n-1]
	dummy := func(nums1, nums2 []int) int {
		ans := 0
		for i := 0; i < n-1; i++ {
			if nums1[i] > nums1[n-1] || nums2[i] > nums2[n-1] {
				if nums2[i] <= nums1[n-1] && nums1[i] <= nums2[n-1] {
					ans++
					nums1[i], nums2[i] = nums2[i], nums1[i]
				} else {
					return -1
				}
			}
		}
		return ans
	}
	noSwap := dummy(nums1, nums2)
	swap := dummy(nums1X, nums2X)
	if noSwap == -1 && swap == -1 {
		return -1
	}
	if noSwap == -1 {
		return swap + 1
	}
	if swap == -1 {
		return noSwap
	}
	return min(noSwap, swap+1)
}
