package main

func canBeValid(s string, locked string) bool {
	n := len(s)
	fixedOpen := []int{}
	stack := []int{}
	for i := 0; i < n; i++ {
		if s[i] == ')' && locked[i] == '1' {
			if len(fixedOpen) > 0 {
				fixedOpen = fixedOpen[:len(fixedOpen)-1]
			} else {
				if len(stack) == 0 {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		} else {
			if s[i] == '(' && locked[i] == '1' {
				fixedOpen = append(fixedOpen, i)
			} else {
				stack = append(stack, i)
			}
		}
	}
	for len(fixedOpen) > 0 && len(stack) > 0 {
		if fixedOpen[len(fixedOpen)-1] > stack[len(stack)-1] {
			return false
		}
		fixedOpen = fixedOpen[:len(fixedOpen)-1]
		stack = stack[:len(stack)-1]
	}
	if len(fixedOpen) > 0 {
		return false
	}
	if len(stack)%2 == 1 {
		return false
	}
	return true
}
