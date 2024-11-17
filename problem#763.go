package main

import "slices"

func partitionLabels(s string) []int {
	mp := map[rune][]int{}
	for i, c := range s {
		if mp[c] == nil {
			mp[c] = []int{i, i}
		} else {
			mp[c][1] = i
		}
	}
	partitions := [][]int{}
	for _, p := range mp {
		partitions = append(partitions, p)
	}
	slices.SortFunc(partitions, func(p1, p2 []int) int {
		if p1[0] != p2[0] {
			return p1[0] - p2[0]
		} else {
			return p1[1] - p2[1]
		}
	})
	ans := []int{}
	w := partitions[0]
	for i := 1; i < len(partitions); i++ {
		p := partitions[i]
		if p[0] < w[1] {
			w[1] = max(w[1], p[1])
		} else {
			ans = append(ans, w[1]-w[0]+1)
			w = p
		}
	}
	ans = append(ans, w[1]-w[0]+1)
	return ans
}
