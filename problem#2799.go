package main

func countCompleteSubarrays(nums []int) int {
	da := map[int]bool{} // distinct array
	for _, n := range nums {
		da[n] = true
	}
	ds := map[int]int{} // distinct subarray
	l, r, n := 0, 0, len(nums)
	ans := 0

	for r < n {
		ds[nums[r]]++
		ll := l
		for len(ds) == len(da) {
			ds[nums[l]]--
			if ds[nums[l]] == 0 {
				delete(ds, nums[l])
			}
			l++
		}
		ans += (n - r) * (l - ll)
		r++
	}

	return ans
}
