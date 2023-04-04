package main
import (
    "fmt"
    "os"
    "bufio"
	"sort"
	"math"
)

var (
    out = bufio.NewWriter(os.Stdout)
    in = bufio.NewReader(os.Stdin)
)

func lowerBound(arr []int, target int) int {
    return sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
}

func sortMapByKey(m map[int]int) []int {
    keys := make([]int, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)
	return keys
}

func main() {
    defer out.Flush()
    var T  = 0
    var n , m int
    fmt.Fscan(in, &T)
	for ; T > 0 ; T -- {
		fmt.Fscan(in, &n , &m)
		var ma = make(map[int]int, n)
		for i := 0 ; i < n ; i ++ {
			var t int
			fmt.Fscan(in , &t)
			ma[t] = ma[t] + 1
		}
		var key = sortMapByKey(ma)
		for i := 0 ; i < m ; i ++ {
			var a , b , c int
			fmt.Fscan(in , &a , &b , &c)
			var v = 4 * a * c
		// 	if(v <= 0) {
		// 		fmt.Fprintln(out , "NO")
		// 	}else {
		// 		var sq = math.Sqrt(float64(v))
		// 		for ; sq * sq >= float64(v) ; sq -- {}
		// 		for ; (sq + 1) * (sq + 1) < float64(v) ; sq ++ {}
		// 		l , r := b - int(sq) , b + int(sq)
		// 		var p = binarySearch(key , 1)
		// 		if p != -1 && p <= r && p >= l{
		// 			fmt.Fprintln(out , "YES")
		// 			fmt.Fprintln(out , p)
		// 		}else {
		// 			fmt.Fprintln(out , "NO")
		// 		}
		// 	}
		// }
	} 
}