package main
import (
	"fmt"
	_"math"
)

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)
	var result int = int(m*n) / 2
	fmt.Println(result)
}