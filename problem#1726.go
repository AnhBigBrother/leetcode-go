package main

func tupleSameProduct(nums []int) int {
	ans := 0
	Dic := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			x := nums[i] * nums[j]
			Dic[x]++
			if Dic[x] > 1 {
				ans += (Dic[x] - 1) * 8
			}
		}
	}
	return ans
}
