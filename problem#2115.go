package main

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ri := map[string][]string{}
	valid := map[string]bool{}
	isRecipe := map[string]bool{}
	visited := map[string]bool{}
	ans := []string{}
	for i := range recipes {
		ri[recipes[i]] = ingredients[i]
		isRecipe[recipes[i]] = true
	}
	for _, s := range supplies {
		valid[s] = true
	}
	var dfs func(r string) bool
	dfs = func(r string) bool {
		val, ok := valid[r]
		if ok || !isRecipe[r] || visited[r] {
			return val
		}
		ret := true
		visited[r] = true
		for _, i := range ri[r] {
			ret = ret && dfs(i)
		}
		valid[r] = ret
		if ret {
			ans = append(ans, r)
		}
		return ret
	}
	for _, r := range recipes {
		dfs(r)
	}
	return ans
}
