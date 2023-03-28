package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
)

var (
    out = bufio.NewWriter(os.Stdout)
    in = bufio.NewReader(os.Stdin)
)

func min(i , j int64) int64 {
	if i <= j {
		return i
	}
	return j
}

func main() {
	var T int;
	defer out.Flush()
	fmt.Fscan(in, &T)
	for ; T > 0 ; T -- {
		var n , k int64
		tot := int64(0)
		fmt.Fscan(in, &n , &k)
		a := make([]int64, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
			tot += a[i]
		}
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		if tot <= k {
			fmt.Fprintln(out,0)
			continue
		}
		ans := tot - k
		for i, s := n-1, a[0]-ans; i > 0; i-- {
			s += a[i]
			var x int64
			y := int64(n - i + 1)
			if s >= 0 {
				x = s / y
			} else {
				x = (s - y + 1) / y
			}
			x = min(x, a[0])
			ans = min(ans, a[0]-x+y-1)
		}
		fmt.Fprintln(out,ans)
	}
}