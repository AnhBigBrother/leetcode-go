package main

import (
	"math"
	"strconv"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, suffix string) int64 {
	count := func(x string) int64 {
		lengthX := len(x)
		lengthSuffix := len(suffix)

		if lengthX < lengthSuffix {
			return 0
		}

		if lengthX == lengthSuffix {
			if x >= suffix {
				return 1
			}
			return 0
		}

		prefixLen := lengthX - lengthSuffix
		var total int64 = 0
		suffixInX := x[prefixLen:]

		for i := range prefixLen {
			currentDigit := int(x[i] - '0')

			if currentDigit > limit {
				total += int64(math.Pow(float64(limit+1), float64(prefixLen-i)))
				return total
			}

			total += int64(currentDigit) * int64(math.Pow(float64(limit+1), float64(prefixLen-i-1)))
		}

		if suffixInX >= suffix {
			total++
		}

		return total
	}

	lowerBound := strconv.FormatInt(start-1, 10)
	upperBound := strconv.FormatInt(finish, 10)

	return count(upperBound) - count(lowerBound)
}
