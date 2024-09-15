package main

import "math"

func findTheLongestSubstring(s string) int {
	ans := 0
	bitMark := make(map[string]int)
	vowels := "aeiou"
	for i, x := range vowels {
		bitMark[string(x)] = int(math.Pow(2, float64(i)))
	}
	prefixXor := make(map[int]int)
	currXOR := 0
	for i, c := range s {
		x := bitMark[string(c)]
		if x > 0 {
			currXOR = currXOR ^ x
		}
		if prefixXor[currXOR] == 0 {
			prefixXor[currXOR] = i + 1
		}
		if currXOR == 0 {
			ans = max(ans, i+1)
		} else {
			ans = max(ans, i+1-prefixXor[currXOR])
		}
	}

	return ans
}
