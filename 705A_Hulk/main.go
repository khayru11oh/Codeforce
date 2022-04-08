package main

import (
	"fmt"
	"strings"
)
func main () {
	var n int
	fmt.Scan(&n)
	var str string = "I hate it "
	if n % 2 == 0 {
		str = strings.Repeat("I hate it I love it ", n/2)
	} else {
		if n == 1 {
			goto N
		} else {
			str = strings.Repeat("I hate it I love it ", n/2) + "I hate it"
		}
	}
	N:
	str = strings.TrimSpace(str)
	fmt.Println(strings.Replace(str, "it ", "that ", n-1))
}