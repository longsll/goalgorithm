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
    var T  = 0
    fmt.Fscan(in, &T)
	for ; T > 0 ; T -- {
		var n , k int
		var s string
		fmt.Fscan(in , &n , &k)
		fmt.Fscan(in , &s)
		f := false
		for i := 0 ; i < len(s) ; i ++ {
			if int(s[i]) - 48 < k {
				fmt.Fprintf(out  ,"%s%d%s\n", s[:i], k, s[i:])
				f = true
				break
			} 
		}
		if f == false {fmt.Fprintf(out ,"%s%d\n" ,s , k)}
	} 
}