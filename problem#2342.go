package main

func maximumSum(nums []int) int {
	DMap := map[int][]int{}
	calSumDigit := func(x int) int {
		s := 0
		for x > 0 {
			s += x % 10
			x /= 10
		}
		return s
	}
	for _, n := range nums {
		s := calSumDigit(n)
		if DMap[s] == nil {
			DMap[s] = []int{}
		}
		if len(DMap[s]) == 0 {
			DMap[s] = append(DMap[s], n)
		} else if len(DMap[s]) == 1 {
			DMap[s] = append(DMap[s], n)
			if DMap[s][0] < DMap[s][1] {
				DMap[s][0], DMap[s][1] = DMap[s][1], DMap[s][0]
			}
		} else {
			if n > DMap[s][0] {
				DMap[s][1], DMap[s][0] = DMap[s][0], n
			} else {
				DMap[s][1] = max(DMap[s][1], n)
			}
		}
	}
	ans := -1
	for _, arr := range DMap {
		if len(arr) > 1 {
			ans = max(ans, arr[0]+arr[1])
		}
	}
	return ans
}
