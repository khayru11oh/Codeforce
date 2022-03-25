package main
import (
	"fmt"
	"strings"
)

func main() {
	var str string = ""
	var upCase, lowCase int = 0, 0
	fmt.Scan(&str)
	for i:=0; i<len(str); i++ {
		if str[i] >='a' && str[i] <= 'z' {
			lowCase++
		} else {
			upCase++
		}
	}
	if upCase > lowCase {
		str = strings.ToUpper(str)
	} else {
		str = strings.ToLower(str)
	}
	fmt.Println(str)
}