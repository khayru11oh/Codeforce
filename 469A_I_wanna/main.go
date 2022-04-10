package main
import (
	"fmt"
)

func inputFunc (mapp map[int]bool, q int) {
	for i:=0; i<q; i++ {
		var j int
		fmt.Scan(&j)
		mapp[j] = true
	}
}

func main () {
	var n int
	fmt.Scan(&n)
	mapp := make(map[int]bool)
	for i:=0; i<2; i++ {
		var q int
		fmt.Scan(&q)
		inputFunc(mapp, q)
	}
	if len(mapp) == n {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}