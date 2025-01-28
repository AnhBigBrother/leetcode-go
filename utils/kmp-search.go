package utils

func FindLPS(s string) []int {
	n := len(s)
	lps := make([]int, n)
	len := 0
	i := 1
	for i < n {
		if s[i] == s[len] {
			len++
			lps[i] = len
			i++
		} else {
			if len == 0 {
				lps[i] = 0
				i++
			} else {
				len = lps[len-1]
			}
		}
	}
	return lps
}

func KmpSearch(str, subStr string) int {
	if len(subStr) > len(str) {
		return -1
	}
	s := subStr + " " + str
	pi := FindLPS(s)
	for i := len(subStr) + 1; i < len(s); i++ {
		if pi[i] == len(subStr) {
			return i - 2*len(subStr)
		}
	}
	return -1
}
