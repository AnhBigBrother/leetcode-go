package main

func wordSubsets(words1 []string, words2 []string) []string {
	m, n := len(words1), len(words2)
	set1 := [][26]uint16{}
	set2 := [][26]uint16{}
	for i := 0; i < m; i++ {
		set1 = append(set1, [26]uint16{})
	}
	for i := 0; i < n; i++ {
		set2 = append(set2, [26]uint16{})
	}
	for i, w := range words1 {
		for _, c := range w {
			set1[i][uint16(c)-97]++
		}
	}
	for i, w := range words2 {
		for _, c := range w {
			set2[i][uint16(c)-97]++
		}
	}
	repesentWord2 := [26]uint16{}
	for i, w := range words2 {
		for _, c := range w {
			repesentWord2[uint16(c)-97] = max(repesentWord2[uint16(c)-97], set2[i][uint16(c)-97])
		}
	}
	ans := []string{}
	for i, s := range set1 {
		flag := true
		for j, r := range repesentWord2 {
			if r > s[j] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, words1[i])
		}
	}
	return ans
}
