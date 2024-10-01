package main

type MyCalendarTwo struct {
	overlapping     [][]int
	non_overlapping [][]int
}

func MyCalendarTwoConstructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (cur *MyCalendarTwo) Book(start int, end int) bool {
	for _, x := range cur.overlapping {
		if x[0] < end && x[1] > start {
			return false
		}
	}
	for _, x := range cur.non_overlapping {
		if x[0] < end && x[1] > start {
			cur.overlapping = append(cur.overlapping, []int{max(x[0], start), min(x[1], end)})
		}
	}
	cur.non_overlapping = append(cur.non_overlapping, []int{start, end})
	return true
}
