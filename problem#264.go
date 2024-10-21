package main

func nthUglyNumber(n int) int {
	arr := make([]int, n)
	arr[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		arr[i] = min(arr[i2]*2, arr[i3]*3, arr[i5]*5)
		if arr[i] == arr[i2]*2 {
			i2++
		}
		if arr[i] == arr[i3]*3 {
			i3++
		}
		if arr[i] == arr[i5]*5 {
			i5++
		}
	}
	return arr[n-1]
}
