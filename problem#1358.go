package main

func numberOfSubstrings1358(s string) int {
	n := len(s)
	a, b, c := make([]int, n+1), make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		a[i] = a[i-1]
		b[i] = b[i-1]
		c[i] = c[i-1]
		if s[i-1] == 'a' {
			a[i]++
		} else if s[i-1] == 'b' {
			b[i]++
		} else {
			c[i]++
		}
	}

	find := func(idx int) int {
		l, r := 0, idx-2
		for l < r {
			m := (l + r + 1) / 2
			if a[idx]-a[m] > 0 &&
				b[idx]-b[m] > 0 &&
				c[idx]-c[m] > 0 {
				l = m
			} else {
				r = m - 1
			}
		}
		return l
	}

	ans := 0
	i := 2
	for ; i <= n; i++ {
		if a[i] > 0 && b[i] > 0 && c[i] > 0 {
			break
		}
	}
	for ; i <= n; i++ {
		j := find(i)
		ans += j + 1
	}

	return ans
}
