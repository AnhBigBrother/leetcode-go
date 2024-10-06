package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	count1 := map[rune]int{}
	count2 := map[rune]int{}
	for _, x := range s1 {
		count1[x]++
	}
	check := func() bool {
		for c := 'a'; c <= 'z'; c++ {
			if count1[c] != count2[c] {
				return false
			}
		}
		return true
	}
	i, j := 0, 0

	for j < len(s2) {
		count2[rune(s2[j])]++
		j++
		if j-i > len(s1) {
			count2[rune(s2[i])]--
			i++
		}
		if check() {
			return true
		}
	}

	return false
}
