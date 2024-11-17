package main

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	res := []int{}
	flag := 0
	l, r := 0, 0
	for r < n {
		if r-1 >= 0 && nums[r] != nums[r-1]+1 {
			flag = r - l
		}
		if r-l+1 >= k {
			if flag == 0 {
				res = append(res, nums[r])
			} else {
				res = append(res, -1)
			}
			if flag > 0 {
				flag--
			}
			l++
		}
		r++
	}
	return res
}
