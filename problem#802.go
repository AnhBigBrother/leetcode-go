package main

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	childs := []int{}
	parent := [][]int{}
	visited := map[int]bool{}
	stack := []int{}

	for i := 0; i < n; i++ {
		childs = append(childs, 0)
		parent = append(parent, []int{})
	}
	for i := range graph {
		if len(graph[i]) == 0 {
			stack = append(stack, i)
			visited[i] = true
		}
		for _, x := range graph[i] {
			childs[i]++
			parent[x] = append(parent[x], i)
		}
	}
	for i := 0; i < len(stack); i++ {
		for _, x := range parent[stack[i]] {
			childs[x]--
			if !visited[x] && childs[x] == 0 {
				visited[x] = true
				stack = append(stack, x)
			}
		}
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if childs[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
