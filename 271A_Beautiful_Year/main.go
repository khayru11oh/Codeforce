package main
import (
	"fmt"
)

func main() {
	var y int
	fmt.Scan(&y)
	for {
		y++
		var a, b, c, d int
		d = y % 10
		c = (y/10)%10
		b = (y/100)%10
		a = y/1000
		if a!=b && a!=c && a!=d && b!=c && b!=d && c!=d {
			fmt.Println(y)
			return
		}
		
	}
}