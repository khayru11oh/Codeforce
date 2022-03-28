package main

import(
	"fmt"
)

func main() {
	var n, summ int = 0, 0
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		var p, q int
		fmt.Scan(&p); fmt.Scan(&q)
		if p+2 <= q {
			summ++
		}
	}
	fmt.Println(summ)
}