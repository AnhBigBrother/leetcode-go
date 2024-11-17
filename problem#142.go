package main

func detectCycle(head *ListNode) *ListNode {
	lm := map[*ListNode]bool{}
	p := head
	for p != nil {
		if lm[p] {
			return p
		}
		lm[p] = true
		p = p.Next
	}
	return nil
}
