package main

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, sum2 := int64(0), int64(0)
	noZero1, noZero2 := true, true
	for _, x := range nums1 {
		sum1 += int64(x)
		if x == 0 {
			sum1++
			noZero1 = false
		}
	}
	for _, x := range nums2 {
		sum2 += int64(x)
		if x == 0 {
			sum2++
			noZero2 = false
		}
	}
	if noZero1 && noZero2 {
		if sum1 == sum2 {
			return sum1
		}
		return -1
	}
	if noZero1 {
		if sum1 >= sum2 {
			return sum1
		}
		return -1
	}
	if noZero2 {
		if sum2 >= sum1 {
			return sum2
		}
		return -1
	}
	return max(sum1, sum2)
}
