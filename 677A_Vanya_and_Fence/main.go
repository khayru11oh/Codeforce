package main
import (
	"fmt"
)

func main() {
	var n, h, summ int = 0, 0, 0
	fmt.Scan(&n)
	fmt.Scan(&h)
	for i:=0; i<n; i++ {
		var helper int
		fmt.Scan(&helper)
		if helper > h {
			summ+=2
		} else {
			summ++
		}
	}
	fmt.Println(summ)
}