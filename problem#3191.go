package main

func minOperations(nums []int) int {
  n := len(nums)
  ans := 0
  
  flip := func (idx int) {
    if nums[idx] == 0 {
      nums[idx] = 1
    } else {
      nums[idx] = 0
    }
  }

  for i:=2; i<n; i++ {
    if nums[i-2] == 0 {
      nums[i-2] = 1
      flip(i-1)
      flip(i)
      ans ++
    }
  }
  if nums[n-2] == 0 || nums[n-1] == 0 {
    return -1
  }
  return ans;
}

