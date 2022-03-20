package main
import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)
	str1, str2 = strings.ToLower(str1), strings.ToLower(str2)
	fmt.Println(strings.Compare(str1, str2))
}