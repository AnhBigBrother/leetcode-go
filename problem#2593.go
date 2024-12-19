package main

import "slices"

type element2593 struct {
	val int
	idx int
}

func findScore(nums []int) int64 {
	pairs := []element2593{}
	marked := map[int]bool{}
	ans := int64(0)
	for i, x := range nums {
		pairs = append(pairs, element2593{val: x, idx: i})
	}
	slices.SortFunc(pairs, func(a, b element2593) int {
		if a.val == b.val {
			if a.idx < b.idx {
				return -1
			}
			return 1
		} else if a.val < b.val {
			return -1
		}
		return 1
	})
	for _, p := range pairs {
		if !marked[p.idx] {
			ans += int64(p.val)
			marked[p.idx] = true
			marked[p.idx+1] = true
			marked[p.idx-1] = true
		}
	}
	return ans
}
