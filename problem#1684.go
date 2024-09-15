package main

func countConsistentStrings(allowed string, words []string) int {
	ans := 0
	hmap := make(map[rune]bool)
	for _, c := range allowed {
		hmap[c] = true
	}
	for _, x := range words {
		flag := true
		for _, c := range x {
			if !hmap[c] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}
