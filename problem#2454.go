package main

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, -1)
	}
	s1, s2 := []int{}, []int{}
	for i := 0; i < n; i++ {
		// s2 and s1 are always decreasing
		for len(s2) > 0 && nums[s2[len(s2)-1]] < nums[i] {
			ans[s2[len(s2)-1]] = nums[i]
			s2 = s2[:len(s2)-1]
		}
		// use tmp stack to ensure s1 is always decreasing
		tmp := []int{}
		for len(s1) > 0 && nums[s1[len(s1)-1]] < nums[i] {
			tmp = append(tmp, s1[len(s1)-1])
			s1 = s1[:len(s1)-1]
		}
		for j := len(tmp) - 1; j >= 0; j-- {
			s2 = append(s2, tmp[j])
		}
		s1 = append(s1, i)
	}
	return ans
}
