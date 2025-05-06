package main

func countSubarrays2444(nums []int, minK int, maxK int) int64 {
	lastIdxMinK, lastIdxMaxK := -1, -1
	startIdx := 0
	var ans int64 = 0
	for i, n := range nums {
		if n < minK || n > maxK {
			lastIdxMinK = -1
			lastIdxMaxK = -1
			startIdx = i + 1
			continue
		}
		if n == minK {
			lastIdxMinK = i
		}
		if n == maxK {
			lastIdxMaxK = i
		}
		idx := min(lastIdxMaxK, lastIdxMinK)
		if idx != -1 {
			ans += int64(idx-startIdx) + 1
		}
	}
	return ans
}
