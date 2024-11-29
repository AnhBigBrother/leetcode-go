package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	arr := []*ListNode{}
	p := head
	for p != nil {
		arr = append(arr, p)
		p = p.Next
	}
	n := len(arr)
	for i := 0; i+k <= n; i += k {
		l, r := i, i+k-1
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	for i := 0; i < n-1; i++ {
		arr[i].Next = arr[i+1]
	}
	arr[n-1].Next = nil
	return arr[0]
}
