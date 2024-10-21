package main

func findKthBit(n int, k int) byte {
	length := []int{0}
	for i := 1; i <= n; i++ {
		length = append(length, length[i-1]*2+1)
	}
	var solve func(n, k int, isReverted bool) byte
	solve = func(n, k int, isReverted bool) byte {
		if n == 1 {
			if isReverted {
				return '1'
			}
			return '0'
		}
		m := (length[n] + 1) / 2
		if k < m {
			return solve(n-1, k, false)
		} else if k > m {
			return solve(n-1, k-m, true)
		}
		if isReverted {
			return '0'
		}
		return '1'
	}
	return solve(n, k, false)
}
