package main

/*
- scan reads user input
- takes a pointer as an argument
- typed data is written to pointer
- retuns number of scanned items (and the error or nil)
*/
import (
	"fmt"
)

func main() {
	var n float64

	fmt.Printf("Please enter a floating point number and press ENTER.\n")

	num, err := fmt.Scan(&n)

	if err == nil {
		fmt.Printf("truncated number: %d", int(n))
	} else {
		fmt.Printf("num: %d", num)
		fmt.Printf("error: %s", err)
	}

}
