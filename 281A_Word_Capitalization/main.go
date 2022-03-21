package main
import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	fmt.Println(strings.ToUpper(string(str[0])) + str[1:])
}