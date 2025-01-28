package main

func countPalindromicSubsequence(s string) int {
	ans := 0
	first := map[rune]int{}
	last := map[rune]int{}
	for i, c := range s {
		if _, ok := first[c]; !ok {
			first[c] = i
		}
		last[c] = i
	}
	for c := range first {
		checked := map[byte]bool{}
		for i := first[c] + 1; i < last[c]; i++ {
			if !checked[s[i]] {
				checked[s[i]] = true
				ans++
			}
		}
	}
	return ans
}
