package main

import (
	"math"
	"slices"
)

func minimizedMaximum(n int, quantities []int) int {
	count := func(maxi int) int {
		res := 0
		for _, x := range quantities {
			res += int(math.Ceil(float64(x) / float64(maxi)))
		}
		return res
	}
	l, r := 1, slices.Max(quantities)
	for l < r {
		m := int(math.Floor(float64(l+r) / 2))
		x := count(m)
		if x > n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
