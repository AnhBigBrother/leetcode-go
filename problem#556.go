package main

import "slices"

func nextGreaterElement3(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	stack := []int{digits[0]}
	for i := 1; i < len(digits); i++ {
		if digits[i] < stack[len(stack)-1] {
			slices.SortFunc(stack, func(x, y int) int {
				if x > y {
					return -1
				}
				return 1
			})
			j := len(stack) - 1
			for stack[j] <= digits[i] {
				j--
			}
			digits[i], stack[j] = stack[j], digits[i]
			for i < len(digits) {
				stack = append(stack, digits[i])
				i++
			}
			ans := 0
			k := 1
			for _, x := range stack {
				ans += k * x
				k *= 10
			}
			if ans > 2147483647 {
				return -1
			}
			return ans
		}
		stack = append(stack, digits[i])
	}
	return -1
}
