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

func min(x , y int)int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
    defer out.Flush()
    var T  = 0
    fmt.Fscan(in, &T)
	for ; T > 0 ; T -- {
		var n , x1 , y1 , x2 , y2 int
		fmt.Fscan(in , &n , &x1 , &y1 , &x2, &y2)
		fmt.Fprintln(out , abs(min(min(x1 - 1, n - x1),min(y1 - 1, n - y1)) - min(min(x2 - 1, n - x2),min(y2 - 1, n - y2))))
	} 
}