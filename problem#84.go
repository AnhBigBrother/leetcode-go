package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{}
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			left[i] = i
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
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
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			right[i] = i
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
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
		ans = max(ans, (right[i]-left[i]+1)*heights[i])
	}
	// fmt.Println(left)
	// fmt.Println(right)
	return ans
}
