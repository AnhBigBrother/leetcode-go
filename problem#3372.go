package main

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	findNeibor := func(ed [][]int) [][]int {
		ans := make([][]int, len(ed)+1)
		for _, e := range ed {
			ans[e[0]] = append(ans[e[0]], e[1])
			ans[e[1]] = append(ans[e[1]], e[0])
		}
		return ans
	}
	nei1 := findNeibor(edges1)
	nei2 := findNeibor(edges2)

	countTargeted1 := func(nei [][]int, x int) int {
		count := 0
		kk := 0
		visited := map[int]bool{x: true}
		next := []int{x}
		for kk < k {
			newNext := []int{}
			for _, a := range next {
				for _, ne := range nei[a] {
					if !visited[ne] {
						visited[ne] = true
						newNext = append(newNext, ne)
						count++
					}
				}
			}
			kk++
			next = newNext
		}
		return count + 1
	}
	countTargeted2 := func(nei [][]int, x int) int {
		count := 0
		kk := 0
		visited := map[int]bool{x: true}
		next := []int{x}
		for kk < k-1 {
			newNext := []int{}
			for _, a := range next {
				for _, ne := range nei[a] {
					if !visited[ne] {
						visited[ne] = true
						newNext = append(newNext, ne)
						count++
					}
				}
			}
			kk++
			next = newNext
		}
		return count
	}

	count1 := []int{}
	count2 := []int{}

	for i := 0; i < len(nei1); i++ {
		count1 = append(count1, countTargeted1(nei1, i))
	}
	for i := 0; i < len(nei2); i++ {
		count2 = append(count2, countTargeted2(nei2, i))
	}

	maxTargeted2 := count2[0]
	for i := 1; i < len(nei2); i++ {
		maxTargeted2 = max(maxTargeted2, count2[i])
	}

	if k == 0 {
		return count1
	}

	ans := []int{}
	for i := 0; i < len(nei1); i++ {
		ans = append(ans, count1[i]+maxTargeted2+1)
	}

	return ans
}
