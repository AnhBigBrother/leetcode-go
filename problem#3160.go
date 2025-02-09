package main

func queryResults(_ int, queries [][]int) []int {
	n := len(queries)
	ans := []int{}
	color_of_balls := map[int]int{}
	count_color := map[int]int{}
	count := 0
	for i := 0; i < n; i++ {
		q := queries[i]
		if color_of_balls[q[0]] > 0 {
			c := color_of_balls[q[0]]
			count_color[c]--
			if count_color[c] == 0 {
				count--
			}
		}
		color_of_balls[q[0]] = q[1]
		count_color[q[1]]++
		if count_color[q[1]] == 1 {
			count++
		}
		ans = append(ans, count)
	}
	return ans
}
