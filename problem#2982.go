package main

func maximumLength(s string) int {
	ans := -1
	count := map[byte][]int{}
	for c := 'a'; c <= 'z'; c++ {
		count[byte(c)] = make([]int, len(s))
	}
	l, r := 0, 0
	for r < len(s) {
		if s[l] != s[r] {
			l = r
		}
		count[s[r]][r-l]++
		r++
	}
	for c := range count {
		sum := 0
		for _, x := range count[c] {
			sum += x
		}
		for i, x := range count[c] {
			if sum >= 3 {
				ans = max(ans, i+1)
			}
			sum -= x
			if sum <= 0 {
				break
			}
		}
	}
	return ans
}
