package main

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	up := [7]int{}
	down := [7]int{}
	check := func(x int) bool {
		for i := range n {
			if tops[i] != x && bottoms[i] != x {
				return false
			}
			if tops[i] == x && bottoms[i] == x {
				continue
			}
			if tops[i] == x {
				up[x]++
			}
			if bottoms[i] == x {
				down[x]++
			}
		}
		return true
	}
	ans := -1
	for x := 1; x <= 6; x++ {
		if check(x) {
			swap := min(up[x], down[x])
			if ans == -1 {
				ans = swap
			} else {
				ans = min(ans, swap)
			}
		}
	}
	return ans
}
