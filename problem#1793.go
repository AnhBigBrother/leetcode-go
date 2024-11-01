package main

func maximumScore(nums []int, k int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] < nums[i] {
			stack = append(stack, i)
			left[i] = i
			continue
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		p := 0
		if len(stack) > 0 {
			p = stack[len(stack)-1] + 1
		}
		left[i] = p
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		if len(stack) == 0 || nums[stack[len(stack)-1]] < nums[i] {
			stack = append(stack, i)
			right[i] = i
			continue
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		p := n - 1
		if len(stack) > 0 {
			p = stack[len(stack)-1] - 1
		}
		right[i] = p
		stack = append(stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		if left[i] <= k && right[i] >= k {
			ans = max(ans, (right[i]-left[i]+1)*nums[i])
		}
	}
	return ans
}
