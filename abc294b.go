package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var v int
			fmt.Scan(&v)
			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(byte(64+v))
			}
		}
		fmt.Println("")
	}
}
