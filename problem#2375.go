package main

func smallestNumber(pattern string) string {
	n := len(pattern)
	ans := make([]byte, n+1)
	visited := map[byte]bool{}

	flag := false

	var backtrack func(cur byte, idx int)
	backtrack = func(cur byte, idx int) {
		ans[idx] = cur
		if idx == n {
			flag = true
			return
		}
		if pattern[idx] == 'I' {
			for i := cur + 1; i <= '9'; i++ {
				if !visited[i] {
					visited[i] = true
					backtrack(i, idx+1)
					if flag {
						return
					}
					visited[i] = false
				}
			}
		} else {
			for i := cur - 1; i >= '1'; i-- {
				if !visited[i] {
					visited[i] = true
					backtrack(i, idx+1)
					if flag {
						return
					}
					visited[i] = false
				}
			}
		}
	}

	for i := byte('1'); i <= '9'; i++ {
		visited[i] = true
		backtrack(i, 0)
		if flag {
			break
		}
		visited[i] = false
	}

	return string(ans)
}
