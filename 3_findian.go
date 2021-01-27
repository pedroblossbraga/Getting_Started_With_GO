/*
===============================
findian.go
Given a string, checks if:
- starts with the letter "i"
- has a letter "a"
- ends with the letter "n"

author: Pedro Blöss Braga
22/01/2021
===============================
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeSpaces(s string, verbose bool) ([]string, int) {
	/*
		Receives a string, and loops through its characters,
		and stores the nonempty values in slice,
		counts quantity of nonletters.
	*/
	N := len(s)
	if verbose == true {
		fmt.Println(N)
	}
	var s_ []string
	var qtd int
	for i := 0; i < N; i++ {

		if string(s[i]) != " " && string(s[i]) != "\n" {
			if verbose == true {
				fmt.Printf("(%s %d)\n", string(s[i]), i)
			}
			s_ = append(s_, string(s[i]))
		} else {
			qtd++
		}
	}
	if verbose == true {
		fmt.Printf("\nquantity of whitespaces or separators:%d\n", qtd)
	}
	return s_, qtd
}

func startsWithI(s []string, verbose bool) bool {
	var i string = "i"
	x := strings.ToLower(string(s[0]))
	if verbose == true {
		fmt.Printf("startswithI: (x: %s, i: %s), result: %t\n", x, i, x == i)
	}
	return x == i
}
func endsWithN(s []string, verbose bool) bool {
	var n string = "n"

	x := strings.ToLower(string(s[len(s)-1]))
	if verbose == true {
		fmt.Printf("lenght of s: %d\n", len(s))
		fmt.Printf("endswithn (x: %s, n: %s) result: %t\n ", x, n, x == n)
	}
	return x == n
}
func hasA(s []string, verbose bool) bool {
	var a string = "a"
	x := strings.ToLower(strings.Join(s, "")) // transform slice into string, to compare
	if verbose == true {
		fmt.Printf("hasA (%s %s) resultado: %t\n", string(x), a, strings.Contains(x, a))
	}
	return strings.Contains(x, a)
}

func main() {
	const verbose = false // change the verbosity of the program

	// scanning for user input of strings (whitespaces allowed)
	fmt.Printf("Please enter a string and press ENTER.\n")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		x := scanner.Text()
		fmt.Printf("Input was: %q\n", x)

		// starts checking
		fmt.Printf("\n\n============== removing spaces =======================\n\n")
		x_slice, qtd := removeSpaces(string(x), verbose) //string => slice of letters (whitespaces removed)
		if verbose == true {
			fmt.Printf("quantity of non letters: %d\n", qtd)
			fmt.Printf("s: %s\n", x_slice)
		}
		fmt.Printf("================= check - starts with I ====================\n\n")
		startsWithI(x_slice, true)
		fmt.Printf("================= check - ends with N ====================\n\n")
		endsWithN(x_slice, true)
		fmt.Printf("================== check - has A ===================\n\n")
		hasA(x_slice, true)
		fmt.Printf("=====================================\n\n")

		if startsWithI(x_slice, verbose) && endsWithN(x_slice, verbose) && hasA(x_slice, verbose) {
			fmt.Printf("\nFound!!\n")
		} else {
			fmt.Printf("\nNot Found.\n")
		}
	}
}
