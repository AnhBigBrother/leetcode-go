package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(arr [](int)) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	idx := 0
	var recurFunc func(node_arr []*TreeNode)
	recurFunc = func(node_arr []*TreeNode) {
		new_node_arr := []*TreeNode{}
		for _, node := range node_arr {
			if node == nil {
				continue
			}
			idx++
			if idx < len(arr) && arr[idx] != -1 {
				node.Left = &TreeNode{Val: arr[idx]}
				new_node_arr = append(new_node_arr, node.Left)
			}
			idx++
			if idx < len(arr) && arr[idx] != -1 {
				node.Right = &TreeNode{Val: arr[idx]}
				new_node_arr = append(new_node_arr, node.Right)
			}
		}
		if len(node_arr) > 0 {
			recurFunc(new_node_arr)
		}
	}
	root := &TreeNode{Val: arr[0]}
	recurFunc([]*TreeNode{root})
	return root
}

func SpreadTree(root *TreeNode) [][]int {
	ans := [][]int{}
	var recurFunc func(root *TreeNode, level int)
	recurFunc = func(root *TreeNode, level int) {
		if len(ans) < level+1 {
			ans = append(ans, []int{})
		}
		if root == nil {
			ans[level] = append(ans[level], -1)
			return
		}
		ans[level] = append(ans[level], root.Val)
		recurFunc(root.Left, level+1)
		recurFunc(root.Right, level+1)
	}
	recurFunc(root, 0)
	fmt.Println("Spreaded Tree:")
	for _, row := range ans {
		fmt.Println(row)
	}
	return ans
}
