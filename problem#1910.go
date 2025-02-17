package main

func removeOccurrences(s string, part string) string {
	n := len(part)
	if len(s) < n {
		return s
	}
	ans := []byte(s[:n])
	for i := n; i < len(s); i++ {
		if len(ans) >= n && string(ans[len(ans)-n:]) == part {
			ans = ans[:len(ans)-n]
		}
		ans = append(ans, s[i])
	}
	if len(ans) >= n && string(ans[len(ans)-n:]) == part {
		ans = ans[:len(ans)-n]
	}
	return string(ans)
}
