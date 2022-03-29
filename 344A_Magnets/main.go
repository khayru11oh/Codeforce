package main

import (
	"fmt"
)

func main () {
	var n, helper, summ int = 0, 0, 0
	fmt.Scan(&n)
	
	for i:=0; i<n; i++ {
		var magnet int
		fmt.Scan(&magnet)
		if i==0 {
			summ++
			helper = magnet
		} else{
			if helper != magnet {
				helper = magnet
				summ++
			}
		}
	}
	fmt.Println(summ)
}