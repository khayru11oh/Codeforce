package main

import (
	"fmt"
)

func main() {
	var a, b, s string
	fmt.Scan(&a, &b)
	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			s += "1"
		} else {
			s += "0"
		}
	}
	fmt.Println(s)
}