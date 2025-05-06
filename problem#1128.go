package main

func numEquivDominoPairs(dominoes [][]int) int {
	count := [100]int{}
	ans := 0
	for _, d := range dominoes {
		x := d[0]
		y := d[1]
		if x > y {
			x, y = d[1], d[0]
		}
		num := x*10 + y
		count[num]++
		if count[num] > 1 {
			ans += count[num] - 1
		}
	}
	return ans
}
