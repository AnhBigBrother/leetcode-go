package main

type item struct {
	start int
	end   int
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	arr := []item{}
	buy, sell := 0, 0
	for i := 1; i < n; i++ {
		for i < n && prices[i-1] >= prices[i] {
			buy = i
			i++
		}
		for i < n && prices[i-1] <= prices[i] {
			sell = i
			i++
		}
		if buy < sell && prices[buy] < prices[sell] {
			arr = append(arr, item{
				start: prices[buy],
				end:   prices[sell],
			})
			buy = i
			sell = i
		}
	}
	for len(arr) > k {
		minProfit := arr[0].end - arr[0].start
		minDuration := 10001
		for i := 1; i < len(arr); i++ {
			minProfit = min(minProfit, arr[i].end-arr[i].start)
			if arr[i].start < arr[i-1].end {
				minDuration = min(minDuration, arr[i-1].end-arr[i].start)
			}
		}
		newArr := []item{}
		flag := true
		if minDuration < minProfit {
			newArr = append(newArr, arr[0])
			for i := 1; i < len(arr); i++ {
				if flag && arr[i-1].end-arr[i].start == minDuration {
					flag = false
					newArr[i-1].end = arr[i].end
					continue
				}
				newArr = append(newArr, arr[i])
			}
		} else {
			for _, a := range arr {
				if flag && a.end-a.start == minProfit {
					flag = false
					continue
				}
				newArr = append(newArr, a)
			}
		}
		arr = newArr
	}
	ans := 0
	for _, x := range arr {
		ans += x.end - x.start
	}
	return ans
}
