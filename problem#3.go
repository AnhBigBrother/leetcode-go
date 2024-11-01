package main

func lengthOfLongestSubstring(s string) int {
	ans := 0
	preIdx := map[rune]int{}
	leng := 0
	for i, c := range s {
		if i == 0 {
			leng = 1
			preIdx[c] = 0
			ans = 1
		} else {
			j, ok := preIdx[c]
			if ok {
				leng = min(leng+1, i-j)
			} else {
				leng++
			}
			preIdx[c] = i
			ans = max(ans, leng)
		}
	}
	return ans
}
