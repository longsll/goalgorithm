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
    var n int
    fmt.Fscan(in, &T)
    for ; T > 0 ; T -- {
		k := 0
		fmt.Fscan(in , &n)
		for i := 0 ; i < n ; i ++ {
			var tmp int
			fmt.Fscan(in ,&tmp)
			k ^= tmp
		}
		if n % 2 != 0 {
				fmt.Fprintln(out , k)	
		}else {
			if k == 0 {
				fmt.Fprintln(out,k + 1)
			}else {	
				k = -1
				fmt.Fprintln(out , k)
			}
		}
		
	} 
}