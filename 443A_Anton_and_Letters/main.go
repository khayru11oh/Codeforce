package main
import (
	"fmt"
	"unicode"
	"bufio"
	"os"
)

func main() {
	consoleRead := bufio.NewReader(os.Stdin)
	input, _ := consoleRead.ReadString('\n')
	mapp := make(map[string]bool)
	for i:=0; i<len(input); i++ {
		if unicode.IsLetter(rune(input[i])) {
			mapp[string(input[i])] = true
		}
	}
	fmt.Println(len(mapp))
}