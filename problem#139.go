package main

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	breakable := []bool{}
	for i := 0; i <= n; i++ {
		breakable = append(breakable, false)
	}
	inDict := make(map[string]bool)
	for _, x := range wordDict {
		inDict[x] = true
	}
	breakable[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if breakable[j] && inDict[s[j:i]] {
				breakable[i] = true
			}
		}
	}
	return breakable[n]
}
