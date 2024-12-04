package main

func canMakeSubsequence(str1 string, str2 string) bool {
	i1, i2 := 0, 0
	for i1 < len(str1) && i2 < len(str2) {
		if str1[i1] == str2[i2] || (str1[i1] == 'z' && str2[i2] == 'a') || str1[i1]+1 == str2[i2] {
			i2++
		}
		i1++
	}
	return i2 == len(str2)
}
