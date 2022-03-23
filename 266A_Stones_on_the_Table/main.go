package main
import (
	"fmt"
)
func main() {
	sum, str, n := 0, "", 0
	fmt.Scan(&n, &str)

	for i:=1; i<len(str); i++ {
		if str[i] == str[i-1] {
			sum++
		}
	}
	fmt.Println(sum)
}