package main

func countGoodNumbers(n int64) int {
	var mod int64 = 1e9 + 7
	var even int64 = (n + 1) / 2
	var odd int64 = n / 2
	var pow func(x, y int64) int64
	pow = func(x, y int64) int64 {
		if y == 0 {
			return 1
		}
		if y%2 == 1 {
			return (x * pow(x, y-1)) % mod
		}
		xx := (x * x) % mod
		y = y / 2
		return pow(xx, y)
	}
	ret := (pow(5, even) * pow(4, odd)) % mod
	return int(ret)
}
