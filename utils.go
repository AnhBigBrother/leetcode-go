package main

func FindGCD(a int, b int) int {
	if a < b {
		c := b
		b = a
		a = c
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		if a%b == 0 {
			return b
		}
		c := b
		b = a % b
		a = c
	}
	return 1
}

type ListNode struct {
	Val  int
	Next *ListNode
}
