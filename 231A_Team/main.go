package main

import (
	"fmt"
)
func main () {
	var num int
	fmt.Scanf("%d", &num)
	var summ int = 0
	for i:=0; i<num; i++ {
		var summ1 int = 0
		for j:=0; j<3; j++ {
			var num1 int
			fmt.Scan(&num1)
			summ1 += num1
		}
		if summ1 >= 2 {
			summ++
		}
	}
	fmt.Println(summ)
}