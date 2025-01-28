package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	next := [][]int{}
	cache := [][]bool{}
	for i := 0; i < numCourses; i++ {
		next = append(next, []int{})
		cache = append(cache, make([]bool, numCourses))
	}
	for _, p := range prerequisites {
		next[p[0]] = append(next[p[0]], p[1])
		cache[p[0]][p[1]] = true
	}
	check := func(u, v int) bool {
		visited := make([]bool, numCourses)
		visited[u] = true
		q := []int{u}
		for i := 0; i < len(q); i++ {
			x := q[i]
			if x == v || cache[x][v] {
				cache[u][v] = true
				return true
			}
			for _, n := range next[x] {
				if !visited[n] {
					q = append(q, n)
					visited[n] = true
				}
			}
		}
		return false
	}
	ans := []bool{}
	for _, q := range queries {
		ans = append(ans, check(q[0], q[1]))
	}
	return ans
}
