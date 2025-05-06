package main

func countInterestingSubarrays(nums []int, mod int, k int) int64 {
	var ans int64 = 0
	cnt := 0
	mpp := map[int]int64{0: 1}

	for _, n := range nums {
		if n%mod == k {
			cnt++
		}

		ans += mpp[(cnt+mod-k)%mod]
		mpp[cnt%mod]++
	}

	return ans
}
