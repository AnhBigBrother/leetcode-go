package main

import "slices"

func minimumOperations(root *TreeNode) int {
	tree := [][]int{}
	var spread func(root *TreeNode, level int)
	spread = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(tree) < level+1 {
			tree = append(tree, []int{})
		}
		tree[level] = append(tree[level], root.Val)
		spread(root.Left, level+1)
		spread(root.Right, level+1)
	}
	spread(root, 0)
	ans := 0
	for _, row := range tree {
		ordered := []int{}
		ordered = append(ordered, row...)
		slices.Sort(ordered)
		rightPos := map[int]int{}
		for i, x := range ordered {
			rightPos[x] = i
		}
		for i := range row {
			for i != rightPos[row[i]] {
				row[i], row[rightPos[row[i]]] = row[rightPos[row[i]]], row[i]
				ans++
			}
		}
	}
	return ans
}
