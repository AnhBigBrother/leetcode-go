package main

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	sum := nums[0]
	for _, x := range nums {
		sum |= x
	}
	ans := 0
	var backtrack func(cur int, idx int)
	backtrack = func(cur int, idx int) {
		if idx >= n {
			return
		}
		cur |= nums[idx]
		if cur == sum {
			ans++
		}
		for i := idx + 1; i < n; i++ {
			backtrack(cur, i)
		}
	}
	for i := range nums {
		backtrack(0, i)
	}
	return ans
}
