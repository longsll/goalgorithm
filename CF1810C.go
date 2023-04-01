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

func min(a int , b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer out.Flush()
	var T, n, c, d int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(in, &n, &c, &d)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
		var per = 0
		var res = c * n + d
		var add = 0
		for i := 0 ; i < n ; i ++ {
			now := a[i]
			if now > per {
				add += d * (now - per - 1)
			}else if now == per {
				add += c;
			}
			res = min(res, add + (n - i - 1) * c);
			per = now;
		} 
		fmt.Println(res)
	}
}
