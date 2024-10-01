package main

func canArrange(arr []int, k int) bool {
	count := make([]int, k)
	for _, x := range arr {
		r := x % k
		if r < 0 {
			r = r + k
		}
		count[r]++
	}
	for i := 1; i < len(count); i++ {
		j := len(count) - i
		if i == j {
			if count[i]%2 != 0 {
				return false
			} else {
				continue
			}
		}
		if count[i] != count[j] {
			return false
		}
	}
	return true
}
