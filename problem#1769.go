package main

func minOperations1769(boxes string) []int {
	n := (len(boxes))
	sum := make([]int, n)
	sum[0] = 0
	count := 0
	if boxes[0] == '1' {
		sum[0] = 1
		count = 1
	}
	pre := sum[0]
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			pre += (i + 1)
			sum[i] += pre
			count++
		} else {
			sum[i] = sum[i-1]
		}
	}
	ans := make([]int, n)
	pre = 0
	for i := 0; i < n; i++ {
		if boxes[i] == '1' {
			pre++
		}
		ans[i] = pre*(i+1) - sum[i] + (sum[n-1] - sum[i] - (count-pre)*(i+1))
	}
	return ans
}
