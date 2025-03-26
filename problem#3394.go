package main

import "slices"

func checkValidCuts(n int, rectangles [][]int) bool {
	x_sections := [][]int{}
	y_sections := [][]int{}
	for _, r := range rectangles {
		x_sections = append(x_sections, []int{r[0], r[2]})
		y_sections = append(y_sections, []int{r[1], r[3]})
	}
	slices.SortFunc(x_sections, func(a, b []int) int {
		if a[0] == b[0] {
			if a[1] < b[1] {
				return -1
			}
			return 0
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	})
	slices.SortFunc(y_sections, func(a, b []int) int {
		if a[0] == b[0] {
			if a[1] < b[1] {
				return -1
			}
			return 0
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	})
	x_count := 0
	x_cur := x_sections[0]
	for _, x := range x_sections {
		if x[0] < x_cur[1] {
			if x[1] > x_cur[1] {
				x_cur[1] = x[1]
			}
		} else {
			x_count++
			x_cur = x
		}
	}
	x_count++
	if x_count >= 3 {
		return true
	}
	y_count := 0
	y_cur := y_sections[0]
	for _, y := range y_sections {
		if y[0] < y_cur[1] {
			if y[1] > y_cur[1] {
				y_cur[1] = y[1]
			}
		} else {
			y_count++
			y_cur = y
		}
	}
	y_count++
	return y_count >= 3
}
