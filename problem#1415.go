package main

func getHappyString(n int, k int) string {
	strs := []string{}

	str := []byte{}
	var backtrack func(cur byte)
	backtrack = func(cur byte) {
		str = append(str, cur)
		if len(str) < n {
			for c := byte('a'); c <= 'c'; c++ {
				if c != cur {
					backtrack(c)
					if len(strs) == k {
						return
					}
					str = str[:len(str)-1]
				}
			}
		} else {
			strs = append(strs, string(str))
		}
	}
	for c := byte('a'); c <= 'c'; c++ {
		backtrack(c)
		str = str[:len(str)-1]
	}

	if len(strs) < k {
		return ""
	}
	return strs[k-1]
}
