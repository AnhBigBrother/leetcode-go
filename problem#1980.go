package main

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	visited := map[string]bool{}
	for _, x := range nums {
		visited[x] = true
	}
	flag := false
	arr := []byte{}

	var backtrack func()
	backtrack = func() {
		if len(arr) == n {
			if !visited[string(arr)] {
				flag = true
			}
			return
		}
		arr = append(arr, '0')
		backtrack()
		if flag {
			return
		}
		arr = arr[:len(arr)-1]
		arr = append(arr, '1')
		backtrack()
		if flag {
			return
		}
		arr = arr[:len(arr)-1]
	}
	backtrack()

	return string(arr)
}
