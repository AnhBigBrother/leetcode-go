package main

import (
	"math/bits"
	"slices"
	"strconv"
)

func minimizeXor(num1 int, num2 int) int {
	n := bits.OnesCount(uint(num2)) - bits.OnesCount(uint(num1))
	bit1 := []byte(strconv.FormatInt(int64(num1), 2))
	if n > 0 {
		ret := []byte{}
		for i := len(bit1) - 1; i >= 0; i-- {
			if bit1[i] == '0' && n > 0 {
				ret = append(ret, '1')
				n--
			} else {
				ret = append(ret, bit1[i])
			}
		}
		for n > 0 {
			ret = append(ret, '1')
			n--
		}
		slices.Reverse(ret)
		x, _ := strconv.ParseUint(string(ret), 2, 64)
		return int(x)
	}
	for i := len(bit1) - 1; i >= 0; i-- {
		if n < 0 && bit1[i] == '1' {
			bit1[i] = '0'
			n++
		}
	}
	x, _ := strconv.ParseUint(string(bit1), 2, 64)
	return int(x)
}
