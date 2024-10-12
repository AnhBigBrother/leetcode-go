package main

func minSwaps(s string) int {
	close := 0
	maxClose := 0
	for _, x := range s {
		if x == ']' {
			close++
			maxClose = max(maxClose, close)
		} else {
			close--
		}
	}
	return (maxClose + 1) / 2
}
