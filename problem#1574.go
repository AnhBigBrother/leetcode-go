package main

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	l, r := 0, n-1
	for r > 0 && arr[r] >= arr[r-1] {
		r--
	}
	res := r
	for l < r {
		for r < n && arr[r] < arr[l] {
			r++
		}
		res = min(res, r-l-1)
		if arr[l+1] < arr[l] {
			break
		}
		l++
	}
	return res
}
