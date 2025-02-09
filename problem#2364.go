package main

func countBadPairs(nums []int) int64 {
	countPair := func(total int) int {
		total -= 0
		return total * (total + 1) / 2
	}
	n := len(nums)
	good_map := map[int]int{}
	for i, x := range nums {
		good_map[x-i]++
	}
	total := countPair(n)
	good := 0
	for _, count := range good_map {
		good += countPair(count)
	}
	return int64(total - good)
}
