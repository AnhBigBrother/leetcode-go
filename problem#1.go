package main

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i, x := range nums {
		if hmap[target-x] > 0 {
			return []int{i, hmap[target-x] - 1}
		} else {
			hmap[x] = i + 1
		}
	}
	return nil
}
