package main

func validStrings(n int) []string {
	arr := [][]string{}
	arr = append(arr, []string{"1", "0"})
	for i := 1; i < n; i++ {
		a := arr[i-1]
		b := []string{}
		for _, x := range a {
			if x[len(x)-1] == '0' {
				b = append(b, x+"1")
			}
			if x[len(x)-1] == '1' {
				b = append(b, x+"0")
				b = append(b, x+"1")
			}
		}
		arr = append(arr, b)
	}
	return arr[n-1]
}
