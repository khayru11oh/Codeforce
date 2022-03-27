package main
import (
	"fmt"
	_"strings"
)

func main() {
	var all, time int = 0, 0
	var str string = ""
	fmt.Scan(&all); fmt.Scan(&time); fmt.Scan(&str)
	for i:=0; i<time; i++ {
		for j:=0; j<len(str)-1; j++ {
			if str[j] == 'B' && str[j+1] == 'G' {
				buf := []rune(str)
				buf[j] = 'G'
				buf[j+1] = 'B'
				str = string(buf)
				j++
			}
		}
	}
	fmt.Println(str)
}