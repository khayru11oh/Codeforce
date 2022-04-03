package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scanln(&n)
	fmt.Println((n / 2) - (n % 2 * n))
}