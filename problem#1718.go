package main

func constructDistancedSequence(n int) []int {
	ans := make([]int, 2*n-1)
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	flag := false
	visited := make([]bool, n+1)
	check := func() bool {
		for _, x := range ans {
			if x == -1 {
				return false
			}
		}
		return true
	}
	var backtrack func(i int)
	backtrack = func(i int) {
		if i >= len(ans) || check() {
			flag = true
			return
		}
		if ans[i] != -1 {
			backtrack(i + 1)
			return
		}
		for x := n; x > 0; x-- {
			if !visited[x] {
				if x == 1 {
					visited[x] = true
					ans[i] = x
					backtrack(i + 1)
					if flag {
						return
					}
					visited[x] = false
					ans[i] = -1
				} else if i+x < len(ans) && ans[i+x] == -1 {
					visited[x] = true
					ans[i] = x
					ans[i+x] = x
					backtrack(i + 1)
					if flag {
						return
					}
					visited[x] = false
					ans[i] = -1
					ans[i+x] = -1
				}
			}
		}
	}
	backtrack(0)

	return ans
}
