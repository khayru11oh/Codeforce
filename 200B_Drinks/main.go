package main
import (
	"fmt"
)

func main () {
	var n int
	var summ float64
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		var num int
		fmt.Scan(&num)
		summ += float64(num)
	}
	fmt.Printf("%.12f", summ/float64(n))
}