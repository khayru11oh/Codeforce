package main
import (
	"fmt"
	"math"
)

func main() {
	var matrix [5][5]int
	var indexOf int = 4 // a middle element's indexes summ(i + j)
	var result int
	for i:=0; i<5; i++ {
		for j:=0; j<5; j++ {
			var temp int
			fmt.Scan(&temp)
			matrix[i][j] = temp
			if matrix[i][j] == 1 {
				result = int(math.Abs(float64(i + j - indexOf)))
			}
		}
	}
	fmt.Println(result)
}