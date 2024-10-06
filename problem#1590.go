package main

func minSubarray(nums []int, p int) int {
	n := len(nums)
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum < p {
		return -1
	}
	k := sum % p
	if k == 0 {
		return 0
	}
	ans := n
	s := 0
	rmap := map[int]int{}
	rmap[k] = 1
	for i := 0; i < n; i++ {
		s += nums[i]
		if rmap[s%p] > 0 {
			ans = min(ans, i+2-rmap[s%p])
		}
		rmap[(s+k)%p] = i + 2
	}
	if ans == n {
		return -1
	}
	return ans
}
