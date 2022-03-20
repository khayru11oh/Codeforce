package main
import (
	"fmt"
	"strings"
	"unicode"
	"sort"
)

func main() {
	var str1 string
	fmt.Scan(&str1)
	f := func (r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r)
	}
	var slice []string = strings.FieldsFunc(str1, f)
	sort.Strings(slice)
	var str2 string = ""
	for i:=0; i<len(slice); i++ {
		str2 = str2 + "+" + slice[i]
	}
	fmt.Println(strings.Trim(str2, "+"))
}