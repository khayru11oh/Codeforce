package main

import (
	"fmt"
)

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	var summ int = 0
	for i:=1; i<=w; i++ {
		summ += k*i
	}
	if (summ - n) < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(summ - n)
	}
}