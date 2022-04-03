package main
import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)
	var slicee []int = make([]int, n)
	for i:=0; i<n; i++ {
		var sliceHelper int
		fmt.Scan(&sliceHelper)
		slicee[i] = sliceHelper
	}
	var j int = 1
	for j<=n {
		for i:=0; i<n; i++ {
			if slicee[i] == j {
				fmt.Print(i+1, "\t")
				break
			}
		}
		j++
	}
}