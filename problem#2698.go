package main

func punishmentNumber(n int) int {
	var check func(x, i int) bool
	check = func(x, i int) bool {
		if x == i {
			return true
		}
		if x < i {
			return false
		}
		y := x % 10
		k := 10
		x /= 10
		for x > 0 {
			if i-y < 0 {
				break
			}
			if check(x, i-y) {
				return true
			}
			y = (x%10)*k + y
			k *= 10
			x /= 10
		}
		return false
	}
	ans := 1
	for i := 9; i <= n; i++ {
		x := i * i
		if check(x, i) {
			ans += x
		}
	}
	return ans
}
