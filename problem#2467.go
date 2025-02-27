package main

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	neighbor := [][]int{}
	for i := 0; i < n; i++ {
		neighbor = append(neighbor, []int{})
	}
	for _, e := range edges {
		neighbor[e[0]] = append(neighbor[e[0]], e[1])
		neighbor[e[1]] = append(neighbor[e[1]], e[0])
	}

	daddy := make([]int, n)
	visited := make([]bool, n)
	var findDaddy func(node int)
	findDaddy = func(node int) {
		for _, x := range neighbor[node] {
			if !visited[x] {
				daddy[x] = node
				visited[x] = true
				findDaddy(x)
			}
		}
	}
	visited[0] = true
	daddy[0] = -1
	findDaddy(0)

	bobPassOnTurn := make([]int, n)
	for i := 0; i < n; i++ {
		bobPassOnTurn[i] = n + 10
	}
	turn := 0
	for bob != -1 {
		bobPassOnTurn[bob] = turn
		bob = daddy[bob]
		turn++
	}

	ans := int(-1e5)
	visited = make([]bool, n)
	visited[0] = true
	var aliceGoDown func(node, point, turn int)
	aliceGoDown = func(node, point, turn int) {
		if turn < bobPassOnTurn[node] {
			point += amount[node]
		} else if turn == bobPassOnTurn[node] {
			point += amount[node] / 2
		}
		if node != 0 && len(neighbor[node]) == 1 {
			ans = max(ans, point)
			return
		}
		for _, x := range neighbor[node] {
			if !visited[x] {
				visited[x] = true
				aliceGoDown(x, point, turn+1)
			}
		}
	}
	aliceGoDown(0, 0, 0)

	return ans
}
