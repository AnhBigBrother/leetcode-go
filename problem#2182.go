package main

func repeatLimitedString(s string, repeatLimit int) string {
	ans := []byte{}
	count := map[byte]int{}
	for _, c := range s {
		count[byte(c)]++
	}

	for c := byte('z'); c >= 'a'; c-- {
		r := 0
		for count[c] > 0 {
			if r == repeatLimit {
				for x := c - 1; x >= 'a'; x-- {
					if count[x] > 0 {
						ans = append(ans, x)
						count[x]--
						r = 0
						break
					}
				}
				if r != 0 {
					break
				}
			}
			ans = append(ans, c)
			count[c]--
			r++
		}
	}

	return string(ans)
}
