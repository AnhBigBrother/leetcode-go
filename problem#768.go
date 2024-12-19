package main

import "slices"

type item768 struct {
	val int
	idx int
}

func maxChunksToSorted(arr []int) int {
	ans := 0
	sorted := []item768{}
	item768s := []item768{}
	for i, x := range arr {
		sorted = append(sorted, item768{
			val: x,
			idx: i,
		})
		item768s = append(item768s, item768{
			val: x,
			idx: i,
		})
	}
	slices.SortFunc[[]item768](sorted, func(a, b item768) int {
		if a.val == b.val {
			if a.idx < b.idx {
				return -1
			}
			return 1
		}
		if a.val < b.val {
			return -1
		}
		return 1
	})
	itemmax768 := item768s[0]
	for i, x := range item768s {
		if itemmax768.val <= x.val {
			itemmax768 = x
		}
		if itemmax768.idx == sorted[i].idx && itemmax768.val == sorted[i].val {
			ans++
		}
	}
	return ans
}
