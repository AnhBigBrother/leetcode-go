package main

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	checkAllZero := func() bool {
		for _, x := range nums {
			if x > 0 {
				return false
			}
		}
		return true
	}
	if checkAllZero() {
		return 0
	}
	check := func(m int) bool {
		diff := make([]int, n+1)
		for i := range m {
			l, r, val := queries[i][0], queries[i][1], queries[i][2]
			diff[l] -= val
			diff[r+1] += val
		}
		w := 0
		for i := range n {
			w += diff[i]
			if nums[i]+w > 0 {
				return false
			}
		}
		return true
	}

	l, r := 0, len(queries)-1
	for l < r {
		m := (l + r) / 2
		if check(m + 1) {
			r = m
		} else {
			l = m + 1
		}
	}
	if !check(r + 1) {
		return -1
	}
	return r + 1
}
