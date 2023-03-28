package main

import "fmt"
const N = 200005

var (
	s []byte
	f [N]int
	t int
	ans int64
)

func main() {
	fmt.Scan(&t)
	for ; t > 0 ; t-- {
		ans = 0
		fmt.Scan(&s)
		len := len(s)
		s = append(s , '0')
		var val , id = 0 , 0
		for i := len ; i >= 1 ; i -- {
			s[i] = s[i - 1]
		}
		for i := 1 ; i <= len ; i ++ {
			if s[i] == '?' {
				f[i] = f[i - 1] + 1
			}else {
				if ((i-id)%2==0&&val==int(s[i])-48)||((i-id)%2==1&&val!=int(s[i])-48)||id==0 {
					f[i]=f[i-1]+1
				} else {
					f[i] = i - id
				} 
				id = i
				val = int(s[i]) - 48
			}
			ans += int64(f[i])
		}
		fmt.Println(ans)
	}
}