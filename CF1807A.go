package main
import "fmt"
func main(){
	var T int
	fmt.Scan(&T)
	for i:= 1 ; i <= T ; i ++ {
		var a , b , c int
		fmt.Scan(&a,&b,&c)
		if a + b == c {
			fmt.Println("+")
		}else{
			fmt.Println("-")
		}
	}
}
