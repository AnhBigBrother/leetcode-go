package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	merge := func(l1, l2 *ListNode) *ListNode {
		if l2 == nil {
			return l1
		}
		if l1 == nil {
			return l2
		}
		p := &ListNode{}
		head := &ListNode{}
		p.Next = head
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				head.Next = l1
				l1 = l1.Next
			} else {
				head.Next = l2
				l2 = l2.Next
			}
			head = head.Next
		}
		if l2 != nil {
			head.Next = l2
		} else if l1 != nil {
			head.Next = l1
		}
		return p.Next.Next
	}
	var solve func(l, r int) *ListNode
	solve = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l+1 == r {
			return merge(lists[l], lists[r])
		}
		m := (l + r) / 2
		l1 := solve(l, m)
		l2 := solve(m+1, r)
		res := merge(l1, l2)
		return res
	}
	return solve(0, len(lists)-1)
}

/*
lists := []*ListNode{MakeList([]int{1, 4, 5}), MakeList([]int{1, 3, 4}), MakeList([]int{2, 6})}
fmt.Println(ListToArr(mergeKLists(lists)))
*/
