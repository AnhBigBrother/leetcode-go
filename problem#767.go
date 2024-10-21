package main

import "slices"

func reorganizeString(s string) string {
	var freq = map[string]int{}
	for _, x := range s {
		freq[string(x)]++
	}
	arr := []string{}
	for k := range freq {
		arr = append(arr, k)
	}
	str := ""
	for len(arr) > 0 {
		slices.SortStableFunc(arr, func(a, b string) int {
			if freq[a] >= freq[b] {
				return 1
			}
			return -1
		})
		if len(str) != 0 && string(str[len(str)-1]) == arr[len(arr)-1] {
			return ""
		}
		str += arr[len(arr)-1]
		freq[arr[len(arr)-1]]--
		if freq[arr[0]] > 0 && string(str[len(str)-1]) != arr[0] {
			str += arr[0]
			freq[arr[0]]--
		}

		for len(arr) > 0 && freq[arr[0]] == 0 {
			arr = arr[1:]
		}
	}
	return str
}
