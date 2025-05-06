package main

func pushDominoes(dominoes string) string {
	do := []byte(dominoes)
	n := len(do)
	q := []int{}
	state := make([]int, n)
	for i := range n {
		if do[i] == 'R' && i+1 < n && do[i+1] == '.' {
			q = append(q, i)
			state[i] = 1
		}
		if do[i] == 'L' && i-1 >= 0 && do[i-1] == '.' {
			q = append(q, i)
			state[i] = 1
		}
	}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if do[i] == '.' {
			continue
		}
		if do[i] == 'R' && i+1 < n {
			if do[i+1] == '.' {
				do[i+1] = 'R'
				q = append(q, i+1)
				state[i+1] = state[i] + 1
			} else if state[i+1] == state[i]+1 {
				do[i+1] = '.'
			}
		}
		if do[i] == 'L' && i-1 >= 0 {
			if do[i-1] == '.' {
				do[i-1] = 'L'
				q = append(q, i-1)
				state[i-1] = state[i] + 1
			} else if state[i-1] == state[i]+1 {
				do[i-1] = '.'
			}
		}
	}
	return string(do)
}
