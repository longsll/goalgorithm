package main

import (
    "fmt"
    "os"
    "bufio"
)

var (
    out = bufio.NewWriter(os.Stdout)
    in = bufio.NewReader(os.Stdin)
)

func main() {
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)
	m := make(map[int]int , n)
	for i := 1 ; i <= n ; i ++ {
		var t int
		fmt.Fscan(in, &t)
		m[t] = m[t] + 1
	}
	var ans int
	for _ , j :=  range m {
		ans += j / 2
	}
	fmt.Fprintln(out,ans)
}