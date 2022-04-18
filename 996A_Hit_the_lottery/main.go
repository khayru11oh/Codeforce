package main
import (
	"fmt"
)
func main() {
	var m, n int = 0, 0
	fmt.Scan(&m)
	n += m / 100; m %= 100
	n += m / 20; m %= 20
	n += m / 10; m %= 10
	n += m / 5; m %= 5
	n += m / 1
	fmt.Println(n)
}