package main

import "strconv"

func recoverFromPreorder(traversal string) *TreeNode {
	parent := map[int]*TreeNode{}
	level := 0
	num := []byte{}
	for i := 0; i < len(traversal); i++ {
		if traversal[i] == '-' {
			level++
		} else {
			num = append(num, traversal[i])
			if i+1 == len(traversal) || traversal[i+1] == '-' {
				nodeVal, _ := strconv.Atoi(string(num))
				num = []byte{}

				node := &TreeNode{
					Val: nodeVal,
				}

				if level != 0 {
					if parent[level].Left == nil {
						parent[level].Left = node
					} else {
						parent[level].Right = node
					}
				}

				parent[level+1] = node
				level = 0
			}
		}
	}

	return parent[1]
}
