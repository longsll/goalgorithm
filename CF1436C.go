package main

import (
	. "fmt"
	"sort"
)

const mod int64 = 1e9 + 7
func A(n, m int) int64 {
	if m > n { return 0 }
	x := int64(1)
	for i := 0; i < m; i++ {
		x = x * int64(n-i) % mod
	}
	return x
}
func main() {
	var n, x, p, l, g int
	Scan(&n, &x, &p)
	sort.Search(n, func(m int) bool {
		if m < p {l++} else if m > p {g++}
		return m > p
	})
	m:=n-1-l-g
	Print(A(x-1,l)*A(n-x,g)%mod*A(m,m)%mod)
}