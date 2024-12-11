package main

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	notSpecial := [][]int{}
	for i := 1; i < n; i++ {
		if i < n && nums[i]%2 == nums[i-1]%2 {
			notSpecial = append(notSpecial, []int{i - 1, i})
		}
	}
	find := func(q []int) bool {
		l, r := 0, len(notSpecial)-1
		for l <= r {
			m := (l + r) / 2
			if notSpecial[m][0] >= q[0] && notSpecial[m][1] <= q[len(q)-1] {
				return false
			} else if notSpecial[m][1] <= q[0] {
				l = m + 1
			} else if notSpecial[m][0] >= q[len(q)-1] {
				r = m - 1
			}
		}
		return true
	}
	ans := []bool{}
	for _, q := range queries {
		ans = append(ans, find(q))
	}
	return ans
}
