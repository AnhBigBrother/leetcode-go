package main

import "fmt"

func shortestPalindrome(s string) string {
	n := len(s)
	chunks := []string{}
	counts := []int{}
	for i := 0; i < n; i++ {
		count := 1
		for i+1 < n && s[i+1] == s[i] {
			count++
			i++
		}
		chunks = append(chunks, string(s[i]))
		counts = append(counts, count)
	}
	m := len(chunks)
	maxLenPalidrome := n
	for j := m - 1; j >= 0; j-- {
		l := 0
		r := j
		memo := maxLenPalidrome
		if chunks[l] == chunks[r] && counts[l] <= counts[r] {
			if counts[l] < counts[r] && l == 0 {
				maxLenPalidrome = memo - (counts[r] - counts[l])
				l++
				r--
			}
			for l <= r {
				if chunks[l] != chunks[r] || counts[l] != counts[r] {
					break
				}
				l++
				r--
			}
			if l >= r {
				break
			}
		}
		maxLenPalidrome = memo - counts[j]
	}
	addFront := ""
	for i := n - 1; i >= maxLenPalidrome; i-- {
		addFront = addFront + string(s[i])
	}
	fmt.Println(chunks)
	fmt.Println(counts)
	fmt.Println(maxLenPalidrome)
	fmt.Println(addFront)
	return addFront + s
}
