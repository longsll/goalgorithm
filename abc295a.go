package main

import "fmt"

func main() {
	m :=  map[string]int{
		"and" : 1,
		"not" : 1,
		"that" : 1,
		"the" : 1,
		"you" : 1,
	}
	var n int
	fmt.Scan(&n)
	f := false
	for i := 0 ; i < n ; i ++ {
		var str string
		fmt.Scan(&str)
		if m[str] == 1{
			f = true
		}
	}
	if f {
		fmt.Print("Yes")
	}else {
		fmt.Print("No")
	}
}