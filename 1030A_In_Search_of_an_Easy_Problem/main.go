package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var T bool = true
	for i:=0; i<n; i++ {
		var easy int
		fmt.Scan(&easy)
		if easy != 0 {
			T = false
		}
	}
	if T {
		fmt.Println("EASY")
	} else {
		fmt.Println("HARD")
	}
}