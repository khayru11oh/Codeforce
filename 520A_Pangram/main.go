package main
import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var str string
	fmt.Scan(&n); fmt.Scan(&str)
	str = strings.ToLower(str)
	mapp := make(map[string]bool)
	for i:=0; i<len(str); i++ {
		mapp[string(str[i])] = true
	}

	if len(mapp) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}