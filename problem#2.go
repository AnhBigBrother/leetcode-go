package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rem := 0
	res := &ListNode{}
	p := res
	for l1 != nil && l2 != nil {
		x := l1.Val + l2.Val + rem
		rem = x / 10
		x = x % 10
		p.Next = &ListNode{Val: x}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		x := l1.Val + rem
		rem = x / 10
		x = x % 10
		p.Next = &ListNode{Val: x}
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		x := l2.Val + rem
		rem = x / 10
		x = x % 10
		p.Next = &ListNode{Val: x}
		p = p.Next
		l2 = l2.Next
	}
	if rem > 0 {
		p.Next = &ListNode{Val: 1}
	}
	return res.Next
}
