package main

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	lis := []int{}
	ans := []int{}
	for _, x := range obstacles {
		if len(lis) == 0 || x >= lis[len(lis)-1] {
			lis = append(lis, x)
			ans = append(ans, len(lis))
			continue
		}
		l, r := 0, len(lis)-1
		for l < r {
			m := (l + r) / 2
			if lis[m] <= x {
				l = m + 1
			} else {
				r = m
			}
		}
		lis[r] = x
		ans = append(ans, r+1)
	}
	return ans
}
