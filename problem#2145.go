package main

func numberOfArrays(diff []int, lower int, upper int) int {
	mi, ma := 0, 0
	x := 0
	for _, d := range diff {
		x += d
		mi = min(x, mi)
		ma = max(x, ma)
	}
	ans := (upper - lower) - (ma - mi) + 1
	if ans < 0 {
		return 0
	}
	return ans
}
