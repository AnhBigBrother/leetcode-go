package main

import "slices"

type Item2818 struct {
	Idx int
	Val int
}

func maximumScore2818(nums []int, k int) int {
	MOD := 1000000007
	n := len(nums)
	e_max := nums[0]
	for _, x := range nums {
		if x > e_max {
			e_max = x
		}
	}
	p_score := make([]int, e_max+1)
	primes := FindPrimeNums(e_max)
	for _, x := range nums {
		if x > 1 && p_score[x] == 0 {
			i := 0
			y := x
			count := 0
			for y >= primes[i] {
				if y%primes[i] == 0 {
					count++
					for y%primes[i] == 0 {
						y /= primes[i]
					}
				} else {
					i++
				}
			}
			p_score[x] = count
		}
	}
	left, right := make([]int, n), make([]int, n)
	l_stack, r_stack := []int{}, []int{}
	for i := range n {
		for len(l_stack) > 0 && p_score[nums[l_stack[len(l_stack)-1]]] < p_score[nums[i]] {
			l_stack = l_stack[:len(l_stack)-1]
		}
		if len(l_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = l_stack[len(l_stack)-1]
		}
		l_stack = append(l_stack, i)
	}
	for i := n - 1; i >= 0; i-- {
		for len(r_stack) > 0 && p_score[nums[r_stack[len(r_stack)-1]]] <= p_score[nums[i]] {
			r_stack = r_stack[:len(r_stack)-1]
		}
		if len(r_stack) == 0 {
			right[i] = n
		} else {
			right[i] = r_stack[len(r_stack)-1]
		}
		r_stack = append(r_stack, i)
	}
	arr := []Item2818{}
	for i := range n {
		arr = append(arr, Item2818{
			Idx: i,
			Val: nums[i],
		})
	}
	slices.SortFunc(arr, func(a, b Item2818) int {
		if a.Val == b.Val {
			if a.Idx < b.Idx {
				return -1
			}
			return 1
		}
		if a.Val < b.Val {
			return 1
		}
		return -1
	})

	ans := 1
	i := 0

	for k > 0 && i < n {
		li, ri := left[arr[i].Idx], right[arr[i].Idx]
		r := (arr[i].Idx - li) * (ri - arr[i].Idx)
		e := min(k, r)
		k -= e
		base := arr[i].Val
		for e > 0 {
			if e%2 == 1 {
				ans = (ans * base) % MOD
			}
			base = (base * base) % MOD
			e /= 2
		}
		i++
	}

	return ans
}
