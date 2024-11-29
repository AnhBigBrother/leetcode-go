package main

func minimumTime(n int, relations [][]int, time []int) int {
	preCourse := make([][]int, n)
	nextCourse := make([][]int, n)
	for _, r := range relations {
		r[0]--
		r[1]--
		preCourse[r[1]] = append(preCourse[r[1]], r[0])
		nextCourse[r[0]] = append(nextCourse[r[0]], r[1])
	}
	timeToFinish := make([]int, n)
	ans := 0
	arr := []int{}
	for i := 0; i < n; i++ {
		if len(nextCourse[i]) == 0 {
			arr = append(arr, i)
		}
	}
	var bfs func(idx int) int
	bfs = func(idx int) int {
		if timeToFinish[idx] > 0 {
			return timeToFinish[idx]
		}
		preTime := 0
		for _, p := range preCourse[idx] {
			preTime = max(preTime, bfs(p))
		}
		timeToFinish[idx] = preTime + time[idx]
		return timeToFinish[idx]
	}
	for _, x := range arr {
		ans = max(ans, bfs(x))
	}
	return ans
}
