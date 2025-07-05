package main

func rob(root *TreeNode) int {
	leaves := []*TreeNode{}
	parent := map[*TreeNode]*TreeNode{}
	max_money := map[*TreeNode]int{}
	q := []*TreeNode{root}
	for i := 0; i < len(q); i++ {
		node := q[i]
		if node.Left != nil {
			parent[node.Left] = node
			q = append(q, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			q = append(q, node.Right)
		}
		if node.Left == nil && node.Right == nil {
			leaves = append(leaves, node)
			max_money[node] = node.Val
		}
	}
	for i := 0; i < len(leaves); i++ {
		node := leaves[i]
		if parent[node] != nil {
			leaves = append(leaves, parent[node])
		}
		l, r := 0, 0
		if node.Left != nil {
			l += max_money[node.Left.Left]
			l += max_money[node.Left.Right]
		}
		if node.Right != nil {
			r += max_money[node.Right.Left]
			r += max_money[node.Right.Right]
		}
		case1 := l + r + node.Val                             // include node
		case2 := max_money[node.Left] + max_money[node.Right] // not include node
		max_money[node] = max(case1, case2)
	}
	return max_money[root]
}
