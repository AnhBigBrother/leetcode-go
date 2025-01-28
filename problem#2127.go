package main

func maximumInvitations(fav []int) int {
	n := len(fav)
	visited := make([]bool, n)
	fav_of := map[int][]int{}
	for i, f := range fav {
		if fav_of[f] == nil {
			fav_of[f] = []int{}
		}
		fav_of[f] = append(fav_of[f], i)
	}

	longestCycle := 0
	pairs := [][]int{}
	for i, f := range fav {
		if !visited[i] && !visited[f] {
			if i == fav[f] {
				pairs = append(pairs, []int{i, f})
				visited[i] = true
				visited[f] = true
			}
		}
	}
	for _, p := range pairs {
		l := 0
		q := []int{p[0]}
		for len(q) > 0 {
			l++
			newQ := []int{}
			for _, i := range q {
				for _, j := range fav_of[i] {
					if !visited[j] {
						visited[j] = true
						newQ = append(newQ, j)
					}
				}
			}
			q = newQ
		}
		q = []int{p[1]}
		for len(q) > 0 {
			l++
			newQ := []int{}
			for _, i := range q {
				for _, j := range fav_of[i] {
					if !visited[j] {
						visited[j] = true
						newQ = append(newQ, j)
					}
				}
			}
			q = newQ
		}
		longestCycle += l
	}

	twoCycleInvitations := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			x := i
			lMap := map[int]int{}
			lMap[x] = 1
			for !visited[x] {
				visited[x] = true
				if lMap[fav[x]] > 0 {
					twoCycleInvitations = max(twoCycleInvitations, lMap[x]-lMap[fav[x]]+1)
					break
				}
				lMap[fav[x]] = lMap[x] + 1
				x = fav[x]
			}
		}
	}

	return max(longestCycle, twoCycleInvitations)
}
