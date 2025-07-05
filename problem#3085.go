package main

import "slices"

func minimumDeletions(word string, k int) int {
	freq := make([]int, 26)
	for _, w := range word {
		freq[w-97]++
	}
	slices.Sort(freq)
	i := 0
	for freq[i] == 0 {
		i++
	}

	ans := len(word)
	del_prev := 0

	for ; i < 26; i++ {
		cur := del_prev
		for j := 25; j > i; j-- {
			if freq[j]-freq[i] <= k {
				break
			}
			cur += freq[j] - freq[i] - k
		}
		del_prev += freq[i]
		ans = min(ans, cur)
	}

	return ans
}
