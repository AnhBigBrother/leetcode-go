package main

func numberOfSubstrings(s string, k int) int {
	n := len(s)
	ans := 0
	l, r := 0, 0
	count := map[byte]int{}
	for r < n {
		count[s[r]]++
		for count[s[r]] >= k {
			ans += n - r
			count[s[l]]--
			l++
		}
		r++
	}
	return ans
}
