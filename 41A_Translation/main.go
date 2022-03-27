package main

import (
	"fmt"
)

func Reverse(str string)(result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var str1, str2 string
	fmt.Scan(&str1); fmt.Scan(&str2)
	if str1 == Reverse(str2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}