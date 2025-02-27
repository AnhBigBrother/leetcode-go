package main

func numOfSubarrays(arr []int) int {
	ans := 0
	odd := 0
	even := 0
	mod := int(1e9 + 7)
	for _, x := range arr {
		if x%2 == 0 {
			even = even + 1
		} else {
			old_even := even
			even = odd
			odd = old_even + 1
		}
		ans += odd
		ans = ans % mod
	}
	return ans
}
