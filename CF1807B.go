package main

import (
	"fmt"
	"sort"
)

func cacl () bool {
	var n int
	fmt.Scan(&n)
	var a  [2][]int
	for i := 0; i < n; i++ {
		v := 0
		fmt.Scan(&v)
		a[v%2] = append(a[v%2], v)
	}
	sort.Ints(a[0])
	sort.Ints(a[1])
	s0 := 0
	for _, v := range a[0] {
		s0 += v
	}
	s1 := 0
	for _, v := range a[1] {
		s1 += v
	}
	return s1 < s0
}

func main(){
	var T int
	fmt.Scan(&T)	 
	for i:=1 ; i <= T ; i ++ {
		if cacl() {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}
}