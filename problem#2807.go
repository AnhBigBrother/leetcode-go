package main

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head

	for p.Next != nil {
		node := &ListNode{Val: FindGCD(p.Val, p.Next.Val), Next: p.Next}
		p.Next = node
		p = node.Next
	}

	return head
}
