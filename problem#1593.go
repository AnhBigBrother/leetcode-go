package main

func maxUniqueSplit(s string) int {
	m := map[string]bool{}
	ans := 0

	var backtrack func(i, count int)
	backtrack = func(i, count int) {
		var str string
		for j := i + 1; j <= len(s); j++ {
			str = s[i:j]
			if !m[str] {
				m[str] = true
				ans = max(ans, count+1)
				backtrack(j, count+1)
				m[str] = false
			}
		}
	}

	backtrack(0, 0)

	return ans
}
