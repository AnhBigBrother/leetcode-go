package main

func closestPrimes(left int, right int) []int {
	sieve := make([]bool, right+1)
	primes := []int{}
	for k := 2; k*k <= right; k++ {
		if !sieve[k] {
			for i := 2; k*i <= right; i++ {
				sieve[k*i] = true
			}
		}
	}
	for i := max(left, 2); i <= right; i++ {
		if !sieve[i] {
			primes = append(primes, i)
		}
	}
	ans := []int{-1, -1}

	for i := len(primes) - 1; i > 0; i-- {
		if ans[1]-ans[0] == 0 ||
			ans[1]-ans[0] >= primes[i]-primes[i-1] {
			ans[1], ans[0] = primes[i], primes[i-1]
		}
	}

	return ans
}
