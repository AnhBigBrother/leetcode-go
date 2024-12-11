package main

func canChange(start string, target string) bool {
	n := len(start)
	i, j := 0, 0
	for i <= n && j <= n {
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}
		if i == n || j == n {
			break
		}
		if start[i] != target[j] {
			return false
		}
		if target[j] == 'L' {
			if i < j {
				return false
			}
		}
		if target[j] == 'R' {
			if i > j {
				return false
			}
		}
		i++
		j++
	}
	return (i == n) && (j == n)
}
