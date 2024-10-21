package main

func maximumSwap(num int) int {
	digits := []int{}
	for num > 0 {
		x := num % 10
		num /= 10
		digits = append(digits, x)
	}
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		idx := i - 1
		for j := i - 1; j >= 0; j-- {
			if digits[j] >= digits[idx] {
				idx = j
			}
		}
		if idx >= 0 && digits[i] < digits[idx] {
			temp := digits[i]
			digits[i] = digits[idx]
			digits[idx] = temp
			break
		}
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans = ans*10 + digits[i]
	}
	return ans
}
