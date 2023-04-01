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
    var T , n int
    fmt.Fscan(in, &T)
	
    for i := 0 ; i < T ; i ++ {
		fmt.Fscan(in, &n)
		f := false
		for j := 1 ; j <= n ; j ++ {
			var k int
			fmt.Fscan(in , &k)
			if j >= k {
				f = true
			}
		}
		if f == true {
			fmt.Println("YES")
		}else {
			fmt.Println("No")
		}
	}
}
    