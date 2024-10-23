package main

import (
	"fmt"
	"slices"
)

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	level_sum := []int64{}

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		for len(level_sum) < level+1 {
			level_sum = append(level_sum, 0)
		}
		level_sum[level] += int64(root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	slices.SortStableFunc(level_sum, func(i, j int64) int {
		if i > j {
			return -1
		}
		return 1
	})
	fmt.Println(level_sum)
	if len(level_sum) < k {
		return -1
	}
	return level_sum[k-1]
}
