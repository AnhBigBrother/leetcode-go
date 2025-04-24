package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	solve := func(str string) string {
		res := strings.Builder{}
		count := 1
		for i := 1; i < len(str); i++ {
			if str[i] != str[i-1] {
				strCount := strconv.Itoa(count)
				res.WriteString(strCount)
				res.WriteByte(str[i-1])
				count = 1
			} else {
				count++
			}
		}
		strCount := strconv.Itoa(count)
		res.WriteString(strCount)
		res.WriteByte(str[len(str)-1])
		return res.String()
	}
	res := "1"
	for i := 1; i < n; i++ {
		res = solve(res)
	}
	return res
}
