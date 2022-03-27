package main
import (
	"fmt"
	_"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var summ int = 0
	var max int = 0

	for i:=0; i<n; i++ {
		var input1, input2 int
		fmt.Scan(&input1); fmt.Scan(&input2)
		if i==0 {
			summ += input1 + input2
		} else if i==n-1 {
			break
		} else {
			summ-=input1
			summ+=input2
		}
		if max < summ {
			max = summ
		}
	}
	fmt.Println(max)
}