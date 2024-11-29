package main

type Node729 struct {
	start int
	end   int
	left  *Node729
	right *Node729
}

func (cur *Node729) insert(start, end int) bool {
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
				cur.right = &Node729{start: start, end: end}
			} else {
				return cur.right.insert(start, end)
			}
		}
	} else {
		if cur.left == nil {
			cur.left = &Node729{start: start, end: end}
		} else {
			return cur.left.insert(start, end)
		}
	}
	return true
}

type MyCalendar struct {
	root *Node729
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{}
}
func (cal *MyCalendar) Book(start int, end int) bool {
	if cal.root == nil {
		cal.root = &Node729{start: start, end: end}
		return true
	}
	return cal.root.insert(start, end)
}
