package main

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	root := [26]int{}
	for i := range root {
		root[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if x == root[x] {
			return x
		}
		return find(root[x])
	}

	union := func(x int, y int) {
		rx := find(x)
		ry := find(y)
		if rx < ry {
			root[ry] = rx
		} else if ry < rx {
			root[rx] = ry
		}
	}

	for i := range s1 {
		x := int(rune(s1[i]) - 97)
		y := int(rune(s2[i]) - 97)
		union(x, y)
	}

	ans := []rune{}
	for _, c := range baseStr {
		x := find(int(rune(c)) - 97)
		ans = append(ans, rune(x+97))
	}

	return string(ans)
}
