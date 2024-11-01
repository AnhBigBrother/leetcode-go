package main

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || root1 != nil && root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	case1 := flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
	case2 := flipEquiv(root1.Right, root2.Left) && flipEquiv(root1.Left, root2.Right)
	return case1 || case2
}

/*
Test :
func main() {
	root1 := MakeTree([]int{1, 2, 3, 4, 5, 6, -1, -1, -1, 7, 8})
	root2 := MakeTree([]int{1, 3, 2, -1, 6, 4, 5, -1, -1, -1, -1, 8, 7})
	fmt.Println(flipEquiv(root1, root2)) // true
}
*/
