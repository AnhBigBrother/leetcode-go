package main

import "slices"

func canPartition(nums []int) bool {

	slices.Sort(nums)

	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 == 1 {
		return false
	}

	s /= 2

	valid := make([]bool, s+1)
	valid[0] = true

	for _, x := range nums {
		if x > s {
			break
		}
		for val := s; val >= x; val-- {
			if valid[val-x] {
				valid[val] = true
			}
		}
	}

	return valid[s]

}
