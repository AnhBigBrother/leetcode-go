package main

func maxDiff(num int) int {
	first_digit := 0
	first_less_than_9 := -1
	digits := map[int]bool{}
	n := num
	for n > 0 {
		d := n % 10
		digits[d] = true
		if d < 9 {
			first_less_than_9 = d
		}
		if n < 10 {
			first_digit = n
		}
		n /= 10
	}

	replace := func(x, y int) int {
		result := 0
		n := num
		k := 1
		for n > 0 {
			z := n % 10
			if z == x {
				result += k * y
			} else {
				result += k * z
			}
			k *= 10
			n /= 10
		}
		return result
	}

	ans := 0
	max_num := num
	if first_less_than_9 != -1 {
		max_num = replace(first_less_than_9, 9)
	}

	for i := range digits {
		if i == first_digit {
			a := replace(i, 1)
			if max_num-a > ans {
				ans = max_num - a
			}
		} else {
			a := replace(i, 0)
			if max_num-a > ans {
				ans = max_num - a
			}
		}
	}

	return ans
}
