package main

func replaceValueInTree(root *TreeNode) *TreeNode {
	sum_at_level := []int{}
	var calLevelSum func(root *TreeNode, level int)
	var recurFunc func(root *TreeNode, level int)

	calLevelSum = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(sum_at_level) < level+1 {
			sum_at_level = append(sum_at_level, 0)
		}
		sum_at_level[level] += root.Val
		calLevelSum(root.Left, level+1)
		calLevelSum(root.Right, level+1)
	}
	recurFunc = func(root *TreeNode, level int) {
		if root == nil || level >= len(sum_at_level) {
			return
		}
		if root.Left != nil && root.Right != nil {
			root.Left.Val = sum_at_level[level+1] - root.Left.Val - root.Right.Val
			root.Right.Val = root.Left.Val
		} else if root.Left != nil {
			root.Left.Val = sum_at_level[level+1] - root.Left.Val
		} else if root.Right != nil {
			root.Right.Val = sum_at_level[level+1] - root.Right.Val
		}
		recurFunc(root.Left, level+1)
		recurFunc(root.Right, level+1)
	}

	root.Val = 0
	calLevelSum(root, 0)
	recurFunc(root, 0)
	return root
}
