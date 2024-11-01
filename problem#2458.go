package main

import "slices"

func treeQueries(root *TreeNode, queries []int) []int {
	height := []int{-1}
	level := [][]int{}
	var dummyFunc func(root *TreeNode, lev int) int

	dummyFunc = func(root *TreeNode, lev int) int {
		if root == nil {
			return -1
		}
		for len(height) <= root.Val {
			height = append(height, 0)
		}
		if len(level) <= lev {
			level = append(level, []int{})
		}
		level[lev] = append(level[lev], root.Val)
		l := dummyFunc(root.Left, lev+1)
		r := dummyFunc(root.Right, lev+1)
		height[root.Val] = max(l, r) + 1
		return height[root.Val]
	}

	dummyFunc(root, 0)
	n := len(height)
	levelOf := make([]int, n)
	levelOf[0] = -1
	for i := 0; i < len(level); i++ {
		slices.SortFunc(level[i], func(x, y int) int {
			if height[x] < height[y] {
				return 1
			}
			return -1
		})
		for _, x := range level[i] {
			levelOf[x] = i
		}
	}

	// fmt.Println(level)
	// fmt.Println(height)
	// fmt.Println(levelOf)

	ans := []int{}
	maxHeight := height[root.Val]

	for _, q := range queries {
		row := level[levelOf[q]]
		if len(row) == 1 {
			ans = append(ans, maxHeight-height[q]-1)
		} else {
			if q == row[0] {
				ans = append(ans, maxHeight-height[q]+height[row[1]])
			} else {
				ans = append(ans, maxHeight)
			}
		}
	}

	return ans
}
