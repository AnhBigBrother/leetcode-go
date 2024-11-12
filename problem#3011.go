package main

func canSortArray(nums []int) bool {
	countSetBitsFunc := func(x int) int {
		count := 0
		for x > 0 {
			count += x & 1
			x >>= 1
		}
		return count
	}
	bitsCounted := map[int]int{}
	for _, x := range nums {
		if bitsCounted[x] == 0 {
			bitsCounted[x] = countSetBitsFunc(x)
		}
	}
	for i := 1; i < len(nums); i++ {
		k := i - 1
		for k >= 0 && nums[k] > nums[k+1] {
			if bitsCounted[nums[k]] != bitsCounted[nums[k+1]] {
				return false
			} else {
				nums[k], nums[k+1] = nums[k+1], nums[k]
				k--
			}
		}
	}
	return true
}
