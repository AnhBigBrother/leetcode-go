package main

type Fwt []int // fenwick tree: https://cp-algorithms.com/data_structures/fenwick.html
func (tree *Fwt) update(idx, val int) {
	idx++
	for idx < len(*tree) {
		(*tree)[idx] += val
		idx += idx & -idx
	}
}
func (tree *Fwt) query(idx int) int {
	ret := 0
	idx++
	for idx > 0 {
		ret += (*tree)[idx]
		idx -= idx & -idx
	}
	return ret
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos2 := make([]int, n)
	for i, x := range nums2 {
		pos2[x] = i
	}

	left := make(Fwt, n+1)
	right := make(Fwt, n+1)

	for i := range n {
		right.update(i, 1)
	}

	var ans int64 = 0

	for _, x := range nums1 {
		pos := pos2[x]

		right.update(pos, -1)

		xChoice := left.query(pos - 1)
		zChoice := right.query(n-1) - right.query(pos)
		ans += int64(xChoice) * int64(zChoice)

		left.update(pos, 1)
	}

	return ans
}
