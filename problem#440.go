package main

func findKthNumber(n int, k int) int {
	ans := 1
	count := func(x int) int {
		res := 0
		y := x + 1
		for x <= n {
			res += min(y, n+1) - x
			x *= 10
			y *= 10
		}
		return res
	}
	for k-1 > 0 {
		steps := count(ans)
		if steps < k {
			ans += 1
			k -= steps
		} else {
			ans *= 10
			k -= 1
		}
	}
	return ans
}
