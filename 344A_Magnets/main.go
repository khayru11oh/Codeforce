package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var summ int = 0
	var strhelper string = ""
	for scanner.Scan() {
		str := scanner.Text()
		if strhelper != str {
			strhelper = str
			summ++
		}
	}
	fmt.Println(summ)
}