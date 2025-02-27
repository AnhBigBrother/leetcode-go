package main

type FindElements struct {
	visited *map[int]bool
}

func NewFindElements(root *TreeNode) FindElements {
	vis := &map[int]bool{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		(*vis)[root.Val] = true
		if root.Left != nil {
			root.Left.Val = root.Val*2 + 1
			dfs(root.Left)
		}
		if root.Right != nil {
			root.Right.Val = root.Val*2 + 2
			dfs(root.Right)
		}
	}
	root.Val = 0
	dfs(root)
	return FindElements{
		visited: vis,
	}
}

func (this *FindElements) Find(target int) bool {
	return (*this.visited)[target]
}
