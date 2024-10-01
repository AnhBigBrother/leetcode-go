package main

type Node struct {
	start int
	end   int
	left  *Node
	right *Node
}

func (cur *Node) insert(start, end int) bool {
	if cur.start == 0 && cur.end == 0 {
		cur.start = start
		cur.end = end
		return true
	}
	if cur.start < end {
		if cur.end > start {
			return false
		} else {
			if cur.right == nil {
				cur.right = &Node{start: start, end: end}
			} else {
				return cur.right.insert(start, end)
			}
		}
	} else {
		if cur.left == nil {
			cur.left = &Node{start: start, end: end}
		} else {
			return cur.left.insert(start, end)
		}
	}
	return true
}

type MyCalendar struct {
	root *Node
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{}
}
func (cal *MyCalendar) Book(start int, end int) bool {
	if cal.root == nil {
		cal.root = &Node{start: start, end: end}
		return true
	}
	return cal.root.insert(start, end)
}
