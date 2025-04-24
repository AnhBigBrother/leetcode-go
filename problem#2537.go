package main

func countGood(nums []int, k int) int64 {
	n := len(nums)
	count := map[int]int{}
	ans := 0
	l, r := 0, 0
	p := 0
	cPair := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cPair[i] = cPair[i-1] + i - 1
	}
	for r < n {
		count[nums[r]]++
		p += cPair[count[nums[r]]]
		if cPair[count[nums[r]]] > 2 {
			p -= cPair[count[nums[r]]-1]
		}
		oldL := l
		for p >= k {
			p -= cPair[count[nums[l]]]
			count[nums[l]]--
			p += cPair[count[nums[l]]]
			l++
		}
		ans += (n - r) * (l - oldL)
		r++
	}
	return int64(ans)
}
