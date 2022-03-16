package main
import (
	"fmt"
	"strings"
)
func main () {
	var n int
	fmt.Scanf("%d", &n)
	var x = 0
	for i:=0; i<n; i++ {
		var operation string
		fmt.Scan(&operation)
		operation = strings.ToLower(operation)
		if strings.Compare(operation, "x++") == 0 || strings.Compare(operation, "++x") == 0 {
			x++
		} else if strings.Compare(operation, "x--") == 0 || strings.Compare(operation, "--x") == 0 {
			x--
		} else {
			continue
		}
	}
	fmt.Println(x)
}