package main

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	right := make([]int, n+1)
	left := make([]int, n+1)
	visited := make([]bool, n+1)

	for i := 1; i < n; i++ {
		right[preorder[i-1]] = preorder[i]
		left[postorder[i]] = postorder[i-1]
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	if n == 1 {
		return root
	}

	var solve func(l, r int, root *TreeNode)
	solve = func(l, r int, root *TreeNode) {
		if visited[l] || visited[r] {
			return
		}
		if l == r {
			visited[l] = true
			node := &TreeNode{Val: l}
			root.Left = node
			solve(right[node.Val], left[node.Val], node)
			return
		}
		visited[l] = true
		visited[r] = true
		leftNode := &TreeNode{Val: l}
		rightNode := &TreeNode{Val: r}
		root.Left = leftNode
		root.Right = rightNode
		solve(right[leftNode.Val], left[leftNode.Val], leftNode)
		solve(right[rightNode.Val], left[rightNode.Val], rightNode)
	}

	visited[root.Val] = true
	visited[0] = true
	solve(preorder[1], postorder[n-2], root)

	return root
}
