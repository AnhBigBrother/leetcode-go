package main

func minBitFlips(start int, goal int) int {
	count := 0
	diff := start ^ goal

	for diff > 0 {
		count += diff & 1
		diff >>= 1
	}

	return count
}
