package main

import (
	"fmt"
)

func main () {
	var summ int = 0
	var k, l, m, n, d int
	fmt.Scanln(&k)
	fmt.Scanln(&l)
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fmt.Scanln(&d)
	for i:=1; i<=d; i++ {
		if i % k != 0 && i % l != 0 && i % m != 0 && i % n != 0 {
			summ++
		}
	}
	fmt.Println(d - summ)
}