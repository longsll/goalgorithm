package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	out = bufio.NewWriter(os.Stdout)
	in  = bufio.NewReader(os.Stdin)
)

func main() {
	defer out.Flush()
	var T, n , c , d int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(in, &n , &c , &d)
		var a , b []int
		var res int
		for i := 0 ; i < n ; i ++ {
			fmt.Fscan(in,&a[i])
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
		
	}
}
