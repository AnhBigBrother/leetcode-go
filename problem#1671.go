package main

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	ans := n
	left := map[int]int{}
	right := map[int]int{}
	for i := 0; i < n; i++ {
		left[i] = n
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				left[i] = min(left[i], left[j]+i-j-1)
			}
		}
		left[i] = min(left[i], i)
	}
	for i := n - 1; i >= 0; i-- {
		right[i] = n
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				right[i] = min(right[i], right[j]+j-i-1)
			}
		}
		right[i] = min(right[i], n-1-i)
	}
	for i := 0; i < n; i++ {
		if left[i] < i && right[i] < n-1-i {
			ans = min(ans, left[i]+right[i])
		}
	}
	return ans
}
