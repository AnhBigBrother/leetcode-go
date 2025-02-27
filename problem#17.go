package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letter := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
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
