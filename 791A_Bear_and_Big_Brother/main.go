package main

import (
	"fmt"
)

func main() {
	var a, b, i int

	fmt.Scan(&a, &b)

	for a<=b {
		a*=3
		b*=2
		i++
	}
	fmt.Println(i)
}