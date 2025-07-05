package main

func maxPathSum(root *TreeNode) int {
	ans := root.Val

	var maxPathDown func(root *TreeNode) int
	maxPathDown = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(0, maxPathDown(root.Left))
		r := max(0, maxPathDown(root.Right))
		ans = max(ans, l+r+root.Val)
		return max(l, r) + root.Val
	}
	maxPathDown(root)

	return ans
}
