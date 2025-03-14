package main

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	aa, ee, ii, oo, uu, kk := make([]int, n+1), make([]int, n+1), make([]int, n+1), make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		aa[i] = aa[i-1]
		ee[i] = ee[i-1]
		ii[i] = ii[i-1]
		oo[i] = oo[i-1]
		uu[i] = uu[i-1]
		kk[i] = kk[i-1]
		switch word[i-1] {
		case 'a':
			aa[i]++
		case 'e':
			ee[i]++
		case 'i':
			ii[i]++
		case 'o':
			oo[i]++
		case 'u':
			uu[i]++
		default:
			kk[i]++
		}
	}

	find := func(left, right int) int {
		l, r := left, right-5
		for l < r {
			m := (l + r + 1) / 2
			if aa[right]-aa[m] > 0 &&
				ee[right]-ee[m] > 0 &&
				ii[right]-ii[m] > 0 &&
				oo[right]-oo[m] > 0 &&
				uu[right]-uu[m] > 0 &&
				kk[right]-kk[m] == k {
				l = m
			} else {
				r = m - 1
			}
		}
		if !(aa[right]-aa[l] > 0 &&
			ee[right]-ee[l] > 0 &&
			ii[right]-ii[l] > 0 &&
			oo[right]-oo[l] > 0 &&
			uu[right]-uu[l] > 0) {
			return 0
		}
		return l - left + 1
	}

	ans := int64(0)
	kkk := 0
	l := 0

	for i := 1; i <= n; i++ {
		if !(word[i-1] == 'a' || word[i-1] == 'e' || word[i-1] == 'i' || word[i-1] == 'o' || word[i-1] == 'u') {
			kkk++
		}
		for kkk > k {
			if !(word[l] == 'a' || word[l] == 'e' || word[l] == 'i' || word[l] == 'o' || word[l] == 'u') {
				kkk--
			}
			l++
		}
		if kkk == k {
			ans += int64(find(l, i))
		}
	}
	return ans
}
