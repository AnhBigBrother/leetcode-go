package main

func minimumIndex(nums []int) int {
	count := map[int]int{}
	for _, x := range nums {
		count[x]++
	}
	do := nums[0]
	for _, x := range nums {
		if count[x] > count[do] {
			do = x
		}
	}
	n := len(nums)
	l := 0
	for i, x := range nums {
		if x == do {
			l++
		}
		r := count[do] - l
		if 2*l > i+1 && 2*r > n-1-i {
			return i
		}
	}
	return -1
}
