package main

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	ans := n
	left := map[int]int{}
	right := map[int]int{}

	stack := []int{0}
	for i := 1; i < n; i++ {
		if nums[stack[len(stack)-1]] < nums[i] {
			left[i] = left[stack[len(stack)-1]] + i - stack[len(stack)-1] - 1
			stack = append(stack, i)
		} else {
			l, r := 0, len(stack)-1
			for l < r {
				m := (l + r) / 2
				if nums[stack[m]] >= nums[i] {
					r = m
				} else {
					l = m + 1
				}
			}
			left[i] = left[stack[l]] + i - stack[l]
			stack[l] = i
		}
	}
	stack = []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		if nums[stack[len(stack)-1]] < nums[i] {
			right[i] = right[stack[len(stack)-1]] + stack[len(stack)-1] - i - 1
			stack = append(stack, i)
		} else {
			l, r := 0, len(stack)-1
			for l < r {
				m := (l + r) / 2
				if nums[stack[m]] >= nums[i] {
					r = m
				} else {
					l = m + 1
				}
			}
			right[i] = right[stack[l]] + stack[l] - i
			stack[l] = i
		}
	}
	for i := 0; i < n; i++ {
		if left[i] < i && right[i] < n-1-i {
			ans = min(ans, left[i]+right[i])
		}
	}
	return ans
}
