package utils

func FindGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return FindGCD(b, a%b)
}

func FindLCM(a, b int) int {
	gcd := FindGCD(a, b)
	return (a / gcd) * b
}

func FindPrimeNums(n int) []int {
	ans := []int{}
	not_primes := map[int]bool{}
	for k := 2; k*k <= n; k++ {
		if !not_primes[k] {
			for i := 2; k*i <= n; i++ {
				not_primes[k*i] = true
			}
		}
	}
	for i := 2; i <= n; i++ {
		if !not_primes[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func MakeArr[T any](leng int, val T) []T {
	arr := []T{}
	for i := 0; i < leng; i++ {
		arr = append(arr, val)
	}
	return arr
}
