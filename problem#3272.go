package main

import "slices"

func countGoodIntegers(n int, k int) int64 {
	var ans int64 = 0
	nn := (n + 1) / 2
	visited := map[int]bool{}

	arrToNum := func(arr []int) int {
		x := 0
		for _, a := range arr {
			x *= 10
			x += a
		}
		return x
	}

	repesent := func(arr []int) int {
		clone := []int{}
		clone = append(clone, arr...)
		slices.SortFunc(clone, func(i, j int) int {
			if i < j {
				return 1
			}
			return -1
		})
		return arrToNum(clone)
	}

	factorial := [11]int{}
	factorial[0] = 1
	for i := 1; i < 11; i++ {
		factorial[i] = factorial[i-1] * i
	}

	cal := func(arr []int) {
		m := n / 2
		for i := range m {
			arr = append(arr, arr[m-1-i])
		}
		x := arrToNum(arr)
		if x%k == 0 {
			r := repesent(arr)
			if visited[r] {
				return
			}
			visited[r] = true
			freq := [10]int{}
			cx := x
			for cx > 0 {
				freq[cx%10]++
				cx /= 10
			}
			x := (n - freq[0]) * factorial[n-1]
			for _, c := range freq {
				x /= factorial[c]
			}
			ans += int64(x)
		}
	}

	var backtrack func(cur []int)
	backtrack = func(cur []int) {
		if len(cur) == nn {
			clone := []int{}
			clone = append(clone, cur...)
			cal(clone)
			return
		}
		for i := 0; i <= 9; i++ {
			cur = append(cur, i)
			backtrack(cur)
			cur = cur[:len(cur)-1]
		}
	}

	for i := 1; i <= 9; i++ {
		backtrack([]int{i})
	}

	return ans
}
