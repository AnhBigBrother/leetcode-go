package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	arr1 := strings.Split(sentence1, " ")
	arr2 := strings.Split(sentence2, " ")
	if len(arr1) > len(arr2) {
		arr3 := arr2
		arr2 = arr1
		arr1 = arr3
	}
	firstCheck := func() bool {
		j := len(arr2) - len(arr1)
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i+j] {
				return false
			}
		}
		return true
	}
	if firstCheck() {
		return true
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
		if i+1 < len(arr1) {
			j := len(arr2) - len(arr1)
			flag := true
			for k := i + 1; k < len(arr1) && k+j < len(arr2); k++ {
				if arr1[k] != arr2[j+k] {
					flag = false
					break
				}
			}
			if flag {
				return true
			}
		} else {
			return true
		}
	}
	return false
}
