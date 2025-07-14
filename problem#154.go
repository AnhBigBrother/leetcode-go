package main

func findMin154(nums []int) int {
	l, r := 0, len(nums)-1
	ans := nums[0]
	for l < r {
		m := (l + r) / 2
		ans = min(ans, nums[l])
		ans = min(ans, nums[r])
		ans = min(ans, nums[m])

		if nums[l] == nums[r] {
			l++
			r--
			continue
		}
		if nums[l] > nums[r] {
			if nums[l] <= nums[m] {
				// [l...m] in ascending order -> findMin(m+1, r)
				l = m + 1
			} else {
				// [m...r] in ascending order -> findMin(l, m-1)
				r = m - 1
			}
		} else {
			// ex: [3,4,9,8,7]
			return ans
		}
	}
	return ans
}
