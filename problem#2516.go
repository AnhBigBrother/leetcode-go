package main

func takeCharacters(s string, k int) int {
	n := len(s)
	strArr := []byte(s)
	for i := 0; i < n; i++ {
		strArr = append(strArr, s[i])
	}
	count := map[byte]int{}
	check := func() bool {
		if count['a'] < k || count['b'] < k || count['c'] < k {
			return false
		}
		return true
	}
	l, r := 0, 0
	for r < n {
		count[strArr[r]]++
		r++
	}
	ans := n + 1
	for l < n && r < len(strArr) {
		for l <= n && check() {
			ans = min(ans, r-l)
			count[strArr[l]]--
			l++
		}
		count[strArr[r]]++
		r++
	}
	if ans > n {
		return -1
	}
	return ans
}
