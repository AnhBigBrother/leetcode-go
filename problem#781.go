package main

import "math"

func numRabbits(answers []int) int {
	count := map[int]int{}
	for _, a := range answers {
		count[a]++
	}
	ans := 0
	for a, c := range count {
		if a == 0 {
			ans += c
			continue
		}
		ans += int(math.Ceil(float64(c)/float64(a+1))) * (a + 1)
	}
	return ans
}
