package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	arr1 := strings.Split(s1, " ")
	arr2 := strings.Split(s2, " ")
	ans := []string{}
	count := make(map[string]int)

	for _, x := range arr1 {
		count[x]++
	}
	for _, x := range arr2 {
		count[x]++
	}
	for _, x := range arr1 {
		if count[x] == 1 {
			ans = append(ans, x)
		}
	}
	for _, x := range arr2 {
		if count[x] == 1 {
			ans = append(ans, x)
		}
	}
	return ans
}
