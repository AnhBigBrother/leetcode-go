package main

func lengthAfterTransformations(s string, t int) int {
	count := [26]int{}
	for _, c := range s {
		count[int(c-'a')]++
	}
	const MOD = 1_000_000_007
	ans := len(s)

	for range t {
		add := [26]int{}
		for i := range 25 {
			add[i+1] += count[i]
		}
		if count[25] > 0 {
			add[0] += count[25]
			add[1] += count[25]
			ans += count[25]
			ans %= MOD
		}
		for i := range 26 {
			count[i] = add[i]
			count[i] %= MOD
		}
	}

	return ans
}
