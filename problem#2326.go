package main

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	r1 := 0
	r2 := m - 1
	c1 := 0
	c2 := n - 1
	var ans [][]int
	for i := 0; i < m; i++ {
		ans = append(ans, []int{})
		for j := 0; j < n; j++ {
			ans[i] = append(ans[i], -1)
		}
	}

	for c1 <= c2 && r1 <= r2 {
		for i := c1; i <= c2; i++ {
			if r1 <= r2 {
				if head == nil {
					ans[r1][i] = -1
				} else {
					ans[r1][i] = head.Val
					head = head.Next
				}
			}
		}
		r1++
		for i := r1; i <= r2; i++ {
			if c1 <= c2 {
				if head == nil {
					ans[i][c2] = -1
				} else {
					ans[i][c2] = head.Val
					head = head.Next
				}
			}
		}
		c2--
		for i := c2; i >= c1; i-- {
			if r1 < r2 {
				if head == nil {
					ans[r2][i] = -1
				} else {
					ans[r2][i] = head.Val
					head = head.Next
				}
			}
		}
		r2--
		for i := r2; i >= r1; i-- {
			if c1 <= c2 {
				if head == nil {
					ans[i][c1] = -1
				} else {
					ans[i][c1] = head.Val
					head = head.Next
				}
			}
		}
		c1++
	}

	return ans
}
