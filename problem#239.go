package main

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	l, r := 0, 0
	ans := []int{}
	for r < k {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		r++
	}
	ans = append(ans, nums[q[0]])
	for r < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if l == q[0] {
			q = q[1:]
		}
		l++
		r++
		ans = append(ans, nums[q[0]])
	}

	return ans
}
