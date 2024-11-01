package main

import (
	"math"
	"slices"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	factoryPositions := []int{}
	for _, f := range factory {
		for f[1] > 0 {
			factoryPositions = append(factoryPositions, f[0])
			f[1]--
		}
	}
	slices.Sort(robot)
	slices.Sort(factoryPositions)
	count_r := len(robot)
	count_f := len(factoryPositions)
	memo := map[int]map[int]int{}
	for i := 0; i <= count_r; i++ {
		memo[i] = map[int]int{}
	}
	for i := 0; i < count_r; i++ {
		memo[i][count_f] = 1e12 // out of factory
	}
	for i := count_r - 1; i >= 0; i-- {
		for j := count_f - 1; j >= 0; j-- {
			assign := math.Abs(float64(robot[i]-factoryPositions[j])) + float64(memo[i+1][j+1])
			skip := memo[i][j+1]
			memo[i][j] = min(int(assign), skip)
		}
	}
	return int64(memo[0][0])
}

/* TEST:
fmt.Println(minimumTotalDistance([]int{1, -1}, [][]int{{-865262624, 6}, {-717666169, 0}, {725929046, 2}, {449443632, 3}, {-912630111, 0}, {270903707, 3}, {-769206598, 2}, {-299780916, 4}, {-159433745, 5}, {-467185764, 3}, {849991650, 7}, {-292158515, 6}, {940410553, 6}, {258278787, 0}, {83034539, 2}, {54441577, 3}, {-235385712, 2}, {75791769, 3}}))
*/
