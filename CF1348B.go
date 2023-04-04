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
	var T, n, k, v int
	for fmt.Fscan(in, &T); T > 0; T-- {
		fmt.Fscan(in, &n, &k)
		a := make([]int, 0, k)
		vis := make([]bool, n+1)
		for i := 0; i < n; i++ {
			if fmt.Fscan(in, &v); !vis[v] {
				vis[v] = true
				a = append(a, v)
			}
		}
		if len(a) > k {
			fmt.Fprintln(out, -1)
			continue
		}
		for i := len(a); i < k; i++ {
			a = append(a, 1)
		}
		fmt.Fprintln(out, len(a)*n)
		for ; n > 0; n-- {
			for _, v := range a {
				fmt.Fprint(out, v, " ")
			}
		}
		fmt.Fprintln(out)
	}   
}