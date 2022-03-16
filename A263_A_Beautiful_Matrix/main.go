package main
import (
	"fmt"
	"math"
)

func main() {
	var matrix [5][5]int
	var indexOfi int = 2
	var indexOfj int = 2
	var result int
	for i:=0; i<5; i++ {
		for j:=0; j<5; j++ {
			var temp int
			fmt.Scan(&temp)
			matrix[i][j] = temp
			if matrix[i][j] == 1 {
				result = int(math.Abs(float64(j - indexOfj)) + math.Abs(float64(i - indexOfi)))
			}
		}
	}
	fmt.Println(result)
}