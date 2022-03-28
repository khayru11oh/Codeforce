package main
import (
	"fmt"
	"strings"
)

func main() {
	var num int64 // pffffff it isn't even used
	var str string
	fmt.Scan(&num); fmt.Scan(&str)
	if strings.Count(str, "A") > strings.Count(str, "D") {
		fmt.Println("Anton")
	} else if strings.Count(str, "A") < strings.Count(str, "D") {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}