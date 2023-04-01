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

func main() {
	defer out.Flush()
	var T, n int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(in, &n)
		if n%2 == 0 {
			fmt.Println("-1")
			continue
		}
		var res []int
		for n > 1 {
			if ((n-1)/2)%2 != 0 {
				res = append(res, 2)
				n = (n - 1) / 2
			} else {
				res = append(res, 1)
				n = (n + 1) / 2
			}
		}
		fmt.Println(len(res))
		for j := len(res) - 1 ; j >= 0 ; j -- {
			fmt.Printf("%d ", res[j])
		}
		fmt.Println("")
	}
}
