package main

import "math"

func assignElements(groups []int, elements []int) []int {
	findDivisors := func(x int) []int {
		ans := []int{1, x}
		visited := map[int]bool{}
		for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
			if x%i == 0 {
				a, b := i, x/i
				if !visited[a] {
					ans = append(ans, a)
					visited[a] = true
				}
				if !visited[b] {
					ans = append(ans, b)
					visited[b] = true
				}
			}
		}
		return ans
	}

	cache := map[int]int{}

	smallestIdx := map[int]int{}
	for i, e := range elements {
		if _, ok := smallestIdx[e]; !ok {
			smallestIdx[e] = i
		}
	}

	ans := []int{}

	for _, g := range groups {
		if idx, ok := cache[g]; ok {
			ans = append(ans, idx)
			continue
		}
		idx := int(1e6)
		divisors := findDivisors(g)
		for _, d := range divisors {
			if i, ok := smallestIdx[d]; ok {
				idx = min(idx, i)
			}
		}
		if idx == 1e6 {
			idx = -1
		}
		ans = append(ans, idx)
		cache[g] = idx
	}

	return ans
}
