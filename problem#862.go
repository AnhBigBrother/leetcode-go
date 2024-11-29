package main

func shortestSubarray(nums []int, k int) int {
	if nums[0] >= k {
		return 1
	}
	n := len(nums)
	queue := []int{0}
	ans := n + 1
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		if nums[i] >= k {
			ans = min(ans, i+1)
		}
		for len(queue) > 0 && nums[i]-nums[queue[0]] >= k {
			ans = min(ans, i-queue[0])
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] <= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	if ans == n+1 {
		return -1
	}
	return ans
}
