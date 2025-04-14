package main

func partitionLabels(s string) []int {
	lastExist := [26]int{}
	for i, c := range s {
		lastExist[int(c)-97] = i
	}
	last := 0
	count := 0
	ans := []int{}
	for i, c := range s {
		count++
		le := lastExist[int(c)-97]
		if le > last {
			last = le
		}
		if i == last {
			ans = append(ans, count)
			count = 0
		}
	}
	return ans
}
