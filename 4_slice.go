/*
=============================================================================================================
- prompts the user 3 integers and stores in a sorted slice in a loop

- create empy integer slice of size (lenght) 3 (array, because it has fixed length)
- during each pass through the loop prompts the user and stores in the slice
- adds the integer to the slice, sorts the slice, prints the contents of the slice in sorted order
- the slice must grow in size to accomodate any number of integers with the user decides to enter
the program should only quit (exiting the loop) when the user enters the "X" chatacter insted of an integer

author: Pedro Blöss Braga
23/01/2021


-> sorting:
	- https://yourbasic.org/golang/how-to-sort-in-go/
	- quick sort, merge sort
=============================================================================================================
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortPrompted(s []int, ind int) {
	/*Sorts the slice only on the numbers that were promped by the user.*/
	s_promped := s[:ind]
	sort.Ints(s_promped)
}
func main() {
	s := make([]int, 3) //slice started with lenght of 3
	var ind int         //index, number of elements in slic
	for {
		fmt.Printf(">>> Enter a number and press ENTER. (or X to finish)\n")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			text := scanner.Text() // user input

			if strings.ToUpper(string(text)) == string("X") {
				break
			} else {
				i, err := strconv.Atoi(text)
				if err == nil {
					if cap(s) < ind+1 {
						aux := make([]int, len(s))
						s = append(s, aux...)
					}
					s[ind] = i
					ind++
				} else {
					fmt.Println("erro: ", err)
				}
			}
			sortPrompted(s, ind)
			fmt.Printf("Input was: %q\nsorted slice: (with extra spaces) %d\nsorted slice (promped numbers):%d\n", text, s, s[:ind])
		}
	}
	fmt.Printf("Finished. slice: %v", s[:ind])
}
