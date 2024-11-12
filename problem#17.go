package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letter := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	ans := []string{}
	char_arr := []byte{}
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(char_arr))
			return
		}
		for _, c := range letter[digits[i]] {
			char_arr = append(char_arr, c)
			backtrack(i + 1)
			char_arr = char_arr[:len(char_arr)-1]
		}
	}
	backtrack(0)
	return ans
}
