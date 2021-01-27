/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”,
  respectively.
 Your program should use Marshal() to create a JSON object from the map,
 and then your program should print the JSON object.

author: Pedro Blöss Braga
25/01/2021

Comments:
- I'm using the bufio scanner in order to enable whitespaces
- if verbosity=true, the program has maximum result printing
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func userInput(s string, verbose bool) string {
	var x string
	fmt.Printf("Enter %s:\n", s)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		x := scanner.Text()
		if verbose == true {
			fmt.Printf("Input was: %q\n", x)
		}
		return string(x)
	}
	return x
}

func main() {
	verbose := false // verbosity of program
	name := userInput("name", verbose)
	addr := userInput("addr", verbose)

	p := make(map[string]string) //creates empy map
	// stores bijection of {key: value}
	p["name"] = name
	p["address"] = addr

	barr, err := json.Marshal(p) // converts map to JSON

	if err == nil {
		fmt.Printf("JSON :%s\n", barr)
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}
