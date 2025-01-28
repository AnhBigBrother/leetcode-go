package main

func SearchFirst(str, subStr string) int {
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

func SearchLast(str, subStr string) int {
	if len(subStr) > len(str) {
		return -1
	}
	s := subStr + " " + str
	pi := FindLPS(s)
	for i := len(s) - 1; i >= len(subStr)+1; i-- {
		if pi[i] == len(subStr) {
			return i - 2*len(subStr)
		}
	}
	return -1
}

func hasMatch(s string, p string) bool {
	idx := 0
	for idx < len(p) {
		if p[idx] == '*' {
			break
		}
		idx++
	}
	first, last := p[:idx], p[idx+1:]
	f, l := SearchFirst(s, first), SearchLast(s, last)
	if len(first) == 0 && l > -1 {
		return true
	}
	if len(last) == 0 && f > -1 {
		return true
	}
	if f != -1 && l != -1 && f+len(first) <= l {
		return true
	}
	return false
}
