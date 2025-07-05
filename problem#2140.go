package main

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	pre := make([]int64, n)
	ans := int64(0)
	for i, q := range questions {
		if i > 0 {
			pre[i] = max(pre[i], pre[i-1])
		}
		j := i + 1 + q[1]
		cur := pre[i] + int64(q[0])
		if j < n {
			pre[j] = max(pre[j], cur)
		}
		ans = max(ans, cur)
	}
	return ans
}
