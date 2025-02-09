package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(arr []int) *ListNode {
	res := &ListNode{}
	p := res
	for _, x := range arr {
		p.Next = &ListNode{Val: x}
		p = p.Next
	}
	return res.Next
}

func LinkedListToArr(head *ListNode) []int {
	ans := []int{}
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}
