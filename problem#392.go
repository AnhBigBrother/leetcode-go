package main

func isSubsequence(s string, t string) bool {
	i1, i2 := 0, 0
	for i1 < len(s) && i2 < len(t) {
		if s[i1] == t[i2] {
			i1++
		}
		i2++
	}
	return i1 == len(s)
}
