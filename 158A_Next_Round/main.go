package main
import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Println("Input contester's scores")
	var scoresSlice []int = make([]int, n)
	for i:=0; i<n; i++ {
		var score int
		fmt.Scan(&score)
		scoresSlice[i] = score
	}
	var kk = scoresSlice[k-1]
	var howmany int
	for i:=0; i<len(scoresSlice); i++ {
		fmt.Println(i, "\t", scoresSlice[i])
		if scoresSlice[i] == 0 {
			continue
		} else {
			if scoresSlice[i] >= kk {
				howmany++
			}
		}
	}
	fmt.Println(howmany)
}