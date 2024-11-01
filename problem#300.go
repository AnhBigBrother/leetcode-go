package main

func lengthOfLIS(nums []int) int {
	binary_dummy := func(arr []int, val, l, r int) {
		for l < r {
			m := (l + r) / 2
			if arr[m] >= val {
				r = m
			} else {
				l = m + 1
			}
		}
		arr[l] = val
	}
	subSequence := []int{}
	for _, x := range nums {
		if len(subSequence) == 0 || subSequence[len(subSequence)-1] < x {
			subSequence = append(subSequence, x)
		} else if subSequence[len(subSequence)-1] > x {
			binary_dummy(subSequence, x, 0, len(subSequence)-1)
		}
	}
	return len(subSequence)
}
