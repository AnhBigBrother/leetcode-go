package main

func maxAbsoluteSum(nums []int) int {
	ans := 0
	pos := 0
	neg := 0
	for _, x := range nums {
		s := pos + x
		if s >= pos {
			pos = s
		} else {
			pos = max(0, s)
		}
		ans = max(ans, pos)
	}
	for _, x := range nums {
		s := neg + x
		if s <= neg {
			neg = s
		} else {
			neg = min(0, s)
		}
		ans = max(ans, -neg)
	}
	return ans
}
