package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	out = bufio.NewWriter(os.Stdout)
	in  = bufio.NewReader(os.Stdin)
)

func change4(t int) byte { //数字1转字符'a'
	return byte(t)
	//fmt.Printf("%c",change4(1))
}

func main() {
	defer out.Flush()
	var T = 0
	var n int
	fmt.Fscan(in, &T)
	for ; T > 0; T-- {
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		var r = 'z'
		var l = 0
		for k, v := range s {
			if int(v) <= int(r) {
				l = k
				r = v
			}
		}
		fmt.Printf("%c%s%s\n", s[l],s[0:l],s[l+1:])
	}
}
