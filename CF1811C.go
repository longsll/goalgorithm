package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	out = bufio.NewWriter(os.Stdout)
	in  = bufio.NewReader(os.Stdin)
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer out.Flush()
	var T = 0
	fmt.Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		n --
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in,&a[i])
		}
		res := make([]int, n+1)
		for i := 1; i < n ; i++ {
			res[i] = min(a[i], a[i-1])
		}
		if a[0] > res[1]{
			res[0] = a[0]
		}
		if a[n - 1] > res[n - 1]{
			res[n] = a[n - 1]
		}
		for i := 0; i <= n; i++ {
			fmt.Fprintf(out ,"%d ", res[i])
		} 
		fmt.Fprintln(out , "")
	}

}
