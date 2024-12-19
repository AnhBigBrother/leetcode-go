package main

func continuousSubarrays(nums []int) int64 {
	ans := 0
	minStack := []int{}
	maxStack := []int{}
	firstSmaller := map[int]int{}
	firstBigger := map[int]int{}

	for i, x := range nums {
		for len(minStack) > 0 && nums[minStack[len(minStack)-1]] >= x {
			minStack = minStack[:len(minStack)-1]
		}
		for len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]] <= x {
			maxStack = maxStack[:len(maxStack)-1]
		}
		if len(minStack) > 0 {
			firstSmaller[i] = minStack[len(minStack)-1]
		} else {
			firstSmaller[i] = -1
		}
		if len(maxStack) > 0 {
			firstBigger[i] = maxStack[len(maxStack)-1]
		} else {
			firstBigger[i] = -1
		}
		minStack = append(minStack, i)
		maxStack = append(maxStack, i)
	}

	l, r := 0, 0
	for r < len(nums) {
		pre_l := l
		fs := firstSmaller[r]
		fb := firstBigger[r]
		for fs > -1 && nums[fs] >= nums[r]-2 {
			fs = firstSmaller[fs]
		}
		for fb > -1 && nums[fb] <= nums[r]+2 {
			fb = firstBigger[fb]
		}
		if l <= fs {
			l = fs + 1
		}
		if l <= fb {
			l = fb + 1
		}
		if pre_l != l {
			ans += (r - pre_l + 1) * (r - pre_l) / 2
			ans -= (r - l + 1) * (r - l) / 2
		}
		r++
	}
	ans += (r - l + 1) * (r - l) / 2
	return int64(ans)
}
