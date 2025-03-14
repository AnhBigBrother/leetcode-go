package main

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	colors = append(colors, colors...)
	count := 0
	l, r := 0, 1
	for ; r < n+k; r++ {
		if r-l == k {
			count++
			l++
		}
		if colors[r] == colors[r-1] {
			l = r
		}
	}
	return count
}
