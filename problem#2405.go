package main

func partitionString(s string) int {
	ans := 0
	visited := [26]bool{}
	for _, c := range s {
		if visited[int(c)-97] {
			ans++
			visited = [26]bool{}
		}
		visited[int(c)-97] = true

	}
	return ans + 1
}
