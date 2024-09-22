package main

func lexicalOrder(n int) []int {
	ans := []int{}
	var recurr func(x int)
	recurr = func(x int) {
		if x > n {
			return
		}
		ans = append(ans, x)
		for i := 0; i < 10; i++ {
			recurr(x*10 + i)
		}
	}
	for i := 1; i < 10; i++ {
		recurr(i)
	}
	return ans
}
