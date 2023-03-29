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
    var s string
    fmt.Fscan(in, &s)
	m := make(map[int]int)
	k := 0
	res := 0
	m[0] = 1
	for i := 0 ; i < len(s) ; i ++ { 
		k ^=  1 << (int(s[i]) - 48)
		res += m[k]
		m[k] = m[k] + 1
	}
	fmt.Fprint(out, res)
}