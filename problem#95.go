package main

func generateTrees(n int) []*TreeNode {
	visited := make([]bool, n+1)

	var recur func(x, min_val, max_val int) []*TreeNode
	recur = func(x, min_val, max_val int) []*TreeNode {
		ans := []*TreeNode{}
		left := []*TreeNode{}
		right := []*TreeNode{}
		for i := min_val; i < x; i++ {
			if !visited[i] {
				visited[i] = true
				left = append(left, recur(i, min_val, x-1)...)
				visited[i] = false
			}
		}
		for i := x + 1; i <= max_val; i++ {
			if !visited[i] {
				visited[i] = true
				right = append(right, recur(i, x+1, max_val)...)
				visited[i] = false
			}
		}
		for _, l := range left {
			for _, r := range right {
				ans = append(ans, &TreeNode{
					Val:   x,
					Left:  l,
					Right: r,
				})
			}
		}
		if len(left) == 0 {
			for _, r := range right {
				ans = append(ans, &TreeNode{
					Val:   x,
					Right: r,
				})
			}
		}
		if len(right) == 0 {
			for _, l := range left {
				ans = append(ans, &TreeNode{
					Val:  x,
					Left: l,
				})
			}
		}
		if len(ans) == 0 {
			return []*TreeNode{{Val: x}}
		}
		return ans
	}

	ans := []*TreeNode{}

	for i := 1; i <= n; i++ {
		visited[i] = true
		ans = append(ans, recur(i, 1, n)...)
		visited[i] = false
	}

	return ans
}
