/*
============================================================================================================================================
Getting Started with Go > Week 4
Final Course Activity: read.go
===================================
- Instructions:

Write a program which reads information from a file and represents it in a slice of structs.
 Assume that there is a text file which contains a series of names.
  Each line of the text file has a first name and a last name, in that order,
   separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
 create a struct which contains the first and last names found in the file.
 Each struct created will be added to a slice, and after all lines have been read from the file,
 your program will have a slice containing one struct for each line in the file.
 After reading all lines from the file, your program should iterate through your slice of structs
 and print the first and last names found in each struct.

-------------------------------
- Example of file "names.txt":
Joseph Silva
	Andrew Souza
Richard Andrade
Pedro Blöss
Pitagoras Wilson
-------------------------------

- author: Pedro Blöß Braga
- 26/01/2021
============================================================================================================================================
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	/*Your program will define a name struct which has two fields,
	fname for the first name, and lname for the last name.*/
	fname string
	lname string
}

func ReadFile(filename string) string {
	f, err := os.Open(filename)

	if err == nil {
		barr := make([]byte, 100)

		nb, err := f.Read(barr)
		if err == nil {
			fmt.Printf("number of bytes: %d\n", nb)
		} else {
			fmt.Printf("error: %s\n", err)
		}
		return string(barr)
	} else {
		fmt.Printf("error: %s\n", err)
		return ""
	}
}

func AssertFieldLenght(s string) string {
	/* Asserting "Each field will be a string of size 20 (characters)."*/
	N := len(s)
	if N < 20 {
		s = s + strings.Repeat(" ", 20-N)
	} else if N > 20 {
		s = string(s[:20])
	}
	return s
}
func GetNumberOfLines(filename string) int {
	_, err := os.Open(filename)
	if err == nil {
		barr := make([]byte, 100)
		if err == nil {
			return len(strings.Split(string(barr), "\n"))
		} else {
			fmt.Printf("err: %s\n", err)
			return 0
		}
	}
	fmt.Printf("err: %s\n", err)
	return 0
}

func GetNamesSlice(filename string) []Person {
	N := GetNumberOfLines(filename)

	f, _ := os.Open(filename)

	reader := bufio.NewReader(f)
	/* Your program will successively read each line of the text file and
	create a struct which contains the first and last names found in the file.*/
	s := make([]Person, 0, N)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		names := strings.Split(string(line), " ")
		firstname, lastname := names[0], names[1]
		s = append(s, Person{fname: firstname, lname: lastname}) //Each struct created will be added to a slice
	}
	return s // slice of structs
}

func main() {
	//filename := "names.txt"
	verbose := false // controls the verbosity of the program

	/*Your program should prompt the user for the name of the text file.*/
	var filename string // will store the name of the file to read (names.txt)
	fmt.Println("Enter the name of the file and press ENTER. (ex: names.txt)")
	fmt.Scanln(&filename) // user input

	if verbose == true {
		dec := "==================="
		f := ReadFile(filename)
		fmt.Printf("%s \n%s \n%s \n", dec, f, dec)
	}

	// get the slice of structs {fname: val, lname: val}
	s := GetNamesSlice(filename) //slice containing one struct for each line in the file

	//iteration through slice of structs
	for index, val := range s {
		fmt.Printf("(line %d) names: %s\n", index, val)
	}
}
