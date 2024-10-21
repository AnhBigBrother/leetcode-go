package main

func parseBoolExpr(expression string) bool {
	if expression == "f" {
		return false
	} else if expression == "t" {
		return true
	}
	op := expression[0]
	expression = expression[2 : len(expression)-1]
	if op == '!' {
		return !parseBoolExpr(expression)
	}
	arr := []string{}
	open := 0
	str := ""
	for _, x := range expression {
		if x == ',' && open == 0 {
			arr = append(arr, str)
			str = ""
			continue
		}
		if x == '(' {
			open++
		} else if x == ')' {
			open--
		}
		str += string(x)
	}
	if len(str) > 0 {
		arr = append(arr, str)
	}
	var ans bool
	if op == '&' {
		ans = true
		for _, x := range arr {
			ans = ans && parseBoolExpr(x)
		}
	} else if op == '|' {
		ans = false
		for _, x := range arr {
			ans = ans || parseBoolExpr(x)
		}
	}
	return ans
}
