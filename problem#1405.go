package main

import "slices"

func longestDiverseString(a int, b int, c int) string {
	freq := map[string]int{"a": a, "b": b, "c": c}
	char := []string{"a", "b", "c"}
	str := ""
	slices.SortStableFunc(char, func(i, j string) int {
		if freq[i] > freq[j] {
			return -1
		}
		return 1
	})
	for freq[char[0]] > 0 {
		if len(str) > 1 && string(str[len(str)-1]) == char[0] && str[len(str)-1] == str[len(str)-2] {
			if freq[char[1]] == 0 {
				return str
			} else {
				str += char[1]
				freq[char[1]]--
			}
		} else {
			str += char[0]
			freq[char[0]]--
		}
		slices.SortStableFunc(char, func(i, j string) int {
			if freq[i] > freq[j] {
				return -1
			}
			return 1
		})
	}
	return str
}
