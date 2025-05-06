package common

import "math"

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
	sieve := make([]bool, n+1)
	for k := 2; k*k <= n; k++ {
		if !sieve[k] {
			// i start from k, no need to start from 2
			// from 2 to k is calculated when k=2, k=3... k=k-1
			for i := k; k*i <= n; i++ {
				sieve[k*i] = true
			}
		}
	}
	for i := 2; i <= n; i++ {
		if !sieve[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func MakeArr[T any](leng int, val T) []T {
	arr := []T{}
	for range leng {
		arr = append(arr, val)
	}
	return arr
}

func FindDivisors(x int) []int {
	ans := []int{1, x}
	visited := map[int]bool{}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			a, b := i, x/i
			if !visited[a] {
				ans = append(ans, a)
				visited[a] = true
			}
			if !visited[b] {
				ans = append(ans, b)
				visited[b] = true
			}
		}
	}
	return ans
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
