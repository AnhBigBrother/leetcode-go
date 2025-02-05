package main

func magnificentSets(n int, edges [][]int) int {
	next := make([][]int, n+1)
	for _, e := range edges {
		if next[e[0]] == nil {
			next[e[0]] = []int{}
		}
		if next[e[1]] == nil {
			next[e[1]] = []int{}
		}
		next[e[0]] = append(next[e[0]], e[1])
		next[e[1]] = append(next[e[1]], e[0])
	}
	visited := make([]bool, n+1)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	get_adj := func(z int) []int {
		level := map[int]int{}
		level[z] = 1
		q := []int{z}
		for i := 0; i < len(q); i++ {
			x := q[i]
			visited[x] = true
			for _, y := range next[x] {
				if !visited[y] {
					if level[y] == 0 {
						level[y] = level[x] + 1
						q = append(q, y)
					} else {
						if abs(level[y]-level[x]) != 1 {
							return nil
						}
					}
				}
			}
		}
		return q
	}
	bfs := func(z int) int {
		ans := 1
		level := map[int]int{}
		vs := map[int]bool{}
		level[z] = 1
		q := []int{z}
		for i := 0; i < len(q); i++ {
			x := q[i]
			vs[x] = true
			ans = max(ans, level[x])
			for _, y := range next[x] {
				if !vs[y] {
					if level[y] == 0 {
						level[y] = level[x] + 1
						q = append(q, y)
					}
				}
			}
		}
		return ans
	}
	ans := 0
	for x := 1; x <= n; x++ {
		if !visited[x] {
			arr := get_adj(x)
			if arr == nil {
				return -1
			}
			a := 1
			for _, b := range arr {
				a = max(a, bfs(b))
			}
			ans += a
		}
	}
	return ans
}
