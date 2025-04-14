package main

func countSymmetricIntegers(low int, high int) int {
	check := func(x int) bool {
		y := x
		n := 0
		for x > 0 {
			n++
			x /= 10
		}
		if n%2 == 1 {
			return false
		}
		n /= 2
		l, r := 0, 0
		for y > 0 {
			if n > 0 {
				r += y % 10
				n--
			} else {
				l += y % 10
			}
			y /= 10
		}
		return l == r
	}
	ans := 0
	for x := low; x <= high; x++ {
		if check(x) {
			ans++
		}
	}
	return ans
}
