package main

import "sort"

type MyCalendarThree struct {
	count map[int]int
}

func MyCalendarThreeConstructor() MyCalendarThree {
	return MyCalendarThree{count: make(map[int]int)}
}

func (cur *MyCalendarThree) Book(start int, end int) int {
	cur.count[start]++
	cur.count[end]--
	times := []int{}
	for i := range cur.count {
		times = append(times, i)
	}
	sort.Ints(times)
	ans := 0
	on_going := 0
	for _, x := range times {
		on_going += cur.count[x]
		ans = max(ans, on_going)
	}
	return ans
}
