package main

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	sum_len := 0
	for p := 0; p < k; p++ {
		stack := []int{}
		for i := p; i < n; i += k {
			if len(stack) == 0 || stack[len(stack)-1] <= arr[i] {
				stack = append(stack, arr[i])
			} else {
				l, r := 0, len(stack)
				for l < r {
					m := (l + r) / 2
					if stack[m] > arr[i] {
						r = m
					} else {
						l = m + 1
					}
				}
				stack[l] = arr[i]
			}
		}
		sum_len += len(stack)
	}
	return n - sum_len
}
