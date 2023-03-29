package main
import (
    "fmt"
    "os"
    "bufio"
)

var (
    out = bufio.NewWriter(os.Stdout)
    in = bufio.NewReader(os.Stdin)
	g [][] byte
)

func main() {
    defer out.Flush()
    var n , m int
    fmt.Fscan(in, &n , &m)
	//
}