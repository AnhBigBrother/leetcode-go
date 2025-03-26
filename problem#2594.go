package main

import "math"

func repairCars(ranks []int, cars int) int64 {
	n := len(ranks)
	max_rank := ranks[0]
	for _, r := range ranks {
		if r > max_rank {
			max_rank = r
		}
	}
	l := int64(1)
	r := int64(math.Pow(float64((cars/n+1)), 2)) * int64(max_rank)

	for l < r {
		m := int64((l + r) / 2)
		count := 0
		for _, r := range ranks {
			count += int(math.Floor(math.Sqrt(float64(m) / float64(r))))
		}
		if count >= cars {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
