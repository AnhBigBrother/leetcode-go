package main

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	fibSubSeqEndAtIdx := [][][]int{}
	for _ = range n {
		fibSubSeqEndAtIdx = append(fibSubSeqEndAtIdx, [][]int{})
	}
	for i := 2; i < n; i++ {
		for j := range i {
			ss := fibSubSeqEndAtIdx[j]
			for _, s := range ss {
				if len(s) >= 2 && (s)[len(s)-1]+(s)[len(s)-2] == arr[i] {
					sub := []int{}
					sub = append(sub, s...)
					sub = append(sub, arr[i])
					fibSubSeqEndAtIdx[i] = append(fibSubSeqEndAtIdx[i], sub)
				}
			}
		}
		l, r := 0, i-1
		for l < r {
			if arr[l]+arr[r] == arr[i] {
				fibSubSeqEndAtIdx[i] = append(fibSubSeqEndAtIdx[i], []int{arr[l], arr[r], arr[i]})
				l++
				r--
			} else if arr[l]+arr[r] < arr[i] {
				l++
			} else {
				r--
			}
		}
	}
	ans := 0
	for _, ss := range fibSubSeqEndAtIdx {
		for _, s := range ss {
			ans = max(ans, len(s))
		}
	}
	return ans
}
