package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func diffWaysToCompute(exp string) []int {
	ans := []int{}
	isNumber := true
	for i := 0; i < len(exp); i++ {
		if !unicode.IsDigit(rune(exp[i])) {
			isNumber = false
			left := diffWaysToCompute(exp[0:i])
			right := diffWaysToCompute(exp[i+1:])
			for _, x := range left {
				for _, y := range right {
					if exp[i] == '+' {
						ans = append(ans, x+y)
					} else if exp[i] == '-' {
						ans = append(ans, x-y)
					} else {
						ans = append(ans, x*y)
					}
				}
			}
		}
	}
	if isNumber {
		x, err := strconv.Atoi(exp)
		if err != nil {
			fmt.Println(err)
			panic("Something went wrong!")
		}
		ans = append(ans, x)
	}
	return ans
}
