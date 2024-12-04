package main

func validArrangement(pairs [][]int) [][]int {
	in := map[int]int{}
	out := map[int][]int{}
	nums := map[int]bool{}
	for _, p := range pairs {
		if out[p[0]] == nil {
			out[p[0]] = []int{}
		}
		out[p[0]] = append(out[p[0]], p[1])
		in[p[1]]++
		nums[p[0]] = true
		nums[p[1]] = true
	}
	start := -1
	for x := range nums {
		if len(out[x]) == in[x]+1 {
			start = x
			break
		}
	}
	if start == -1 {
		for x := range nums {
			start = x
			break
		}
	}

	reverseAns := [][]int{}

	// Post-order dfs:
	// If a node is in a circle, we will reach that node again in the end of that circle
	// Suppose we have circle path at node curr, there are 2 cases: the path which not have circle (the end path) is at first or not
	// If the end path is encounted first => the end path will allway add first
	// If the end path is not encounted first => all paths encounted before have circle => curr node will reach it selft before in each circle path. With post-order dfs, we will continue walk to the next before add current node, so all the circle path before will not add till the last path was added => the end path still at first.
	var dfs func(curr int)
	dfs = func(curr int) {
		for len(out[curr]) > 0 {
			x := out[curr][0]
			out[curr] = out[curr][1:]
			dfs(x)
			reverseAns = append(reverseAns, []int{curr, x})
		}
	}

	dfs(start)

	ans := [][]int{}
	for i := len(reverseAns) - 1; i >= 0; i-- {
		ans = append(ans, reverseAns[i])
	}

	return ans
}
