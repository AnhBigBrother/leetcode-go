package main

func flatten(root *TreeNode) {
	var flatt func(root *TreeNode) (*TreeNode, *TreeNode)
	flatt = func(root *TreeNode) (*TreeNode, *TreeNode) {
		if root == nil {
			return nil, nil
		}
		l, last_l := flatt(root.Left)
		r, last_r := flatt(root.Right)
		root.Left = nil
		if l != nil {
			root.Right = l
			if r != nil {
				last_l.Right = r
				return root, last_r
			}
			return root, last_l
		}
		if r != nil {
			return root, last_r
		}
		return root, root
	}
	flatt(root)
}
