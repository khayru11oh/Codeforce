package main
import (
	"fmt"
)
func main() {
	var x, s  int =  0, 0
	fmt.Scan(&x)
	for i:=5; i>0; i-- {
		s += int(x / i)
		x %= i
	}
	fmt.Println(s)
}