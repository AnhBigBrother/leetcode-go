package main

func partition(ss string) [][]string {
	ans := [][]string{}
	charArr := []byte(ss)
	isPalindrome := func(s []byte) bool {
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	var backtrack func(s []byte, parti []string)
	backtrack = func(s []byte, parti []string) {
		if len(s) == 0 {
			x := []string{}
			x = append(x, parti...)
			ans = append(ans, x)
		}
		for i := 1; i <= len(s); i++ {
			sss := s[:i]
			if isPalindrome(sss) {
				parti = append(parti, string(sss))
				backtrack(s[i:], parti)
				parti = parti[:len(parti)-1]
			}
		}
	}
	backtrack(charArr, []string{})
	return ans
}
