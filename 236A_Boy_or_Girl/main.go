package main
import (
	"fmt"
	"strings"
)
func main() {
	var str, strHelper string
	fmt.Scan(&str)
	for i:=0; i<len(str); i++ {
		if strings.Count(strHelper, string(str[i])) == 0 {
			strHelper += string(str[i])
		} else {
			continue
		}
	}
	if len(strHelper) % 2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}