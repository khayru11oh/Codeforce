package main

import (
	"fmt"
)
func main() {
	var num int64
	var summ int = 0
	fmt.Scan(&num)
	for {
		if num % 10 == 7 || num % 10 == 4 {
			summ++
			num /= 10
		} else {
			num /= 10
		}
		if num <= 0 {
			break
		}
	}
	if summ == 4 || summ == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}