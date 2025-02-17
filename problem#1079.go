package main

import "slices"

func numTilePossibilities(tiles string) int {
	n := len(tiles)
	t := []byte(tiles)
	slices.Sort(t)
	visited := make([]bool, n)
	count := 0

	var backtrack func()
	backtrack = func() {
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if i > 0 && t[i] == t[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			count++
			backtrack()
			visited[i] = false
		}
	}

	backtrack()

	return count
}
