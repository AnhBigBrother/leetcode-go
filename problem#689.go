package main

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	sums := []int{}
	s := 0
	for i := 0; i < k; i++ {
		s += nums[i]
	}
	for i := k; i < n; i++ {
		sums = append(sums, s)
		s -= nums[i-k]
		s += nums[i]
	}
	sums = append(sums, s)

	left, right := make([]int, n-k+1), make([]int, n-k+1)
	left[0], right[n-k] = 0, n-k
	for i := 1; i <= n-k; i++ {
		if sums[left[i-1]] < sums[i] {
			left[i] = i
		} else {
			left[i] = left[i-1]
		}
	}
	for i := n - k - 1; i >= 0; i-- {
		if sums[right[i+1]] <= sums[i] {
			right[i] = i
		} else {
			right[i] = right[i+1]
		}
	}
	maxSum := 0
	ans := []int{0, 0, 0}
	for i := k; i+k <= n-k; i++ {
		l, r := left[i-k], right[i+k]
		if maxSum < sums[l]+sums[i]+sums[r] {
			maxSum = sums[l] + sums[i] + sums[r]
			ans[0] = l
			ans[1] = i
			ans[2] = r
		}
	}
	return ans
}
