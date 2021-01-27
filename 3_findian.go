package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
func startsWithI(s string) bool {
	var i string = "i"
	//s_ := strings.Replace(s, " ", "", -1)
	x := strings.ToLower(string(s[0]))
	fmt.Printf("startswithI: (x: %s, i: %s)", x, i)
	return x == i
}
*/
func startsWithI(s []string) bool {
	var i string = "i"
	fmt.Printf("\n====================================\n")
	//s_ := strings.Replace(s, " ", "", -1)
	x := strings.ToLower(string(s[0]))
	fmt.Printf("startswithI: (x: %s, i: %s)\n", x, i)
	fmt.Printf("\n====================================\n")
	return x == i
}

/*
func endsWithN(s string) bool {
	var n string = "n"
	//s_ := string(strings.Replace(s, " ", "", -1))

		fmt.Printf("\n====================================\n")
		fmt.Printf("lenght of s: %d", len(s))
		for i := 0; i < len(s)-2; {
			fmt.Printf("(i: %d, char: %s)\n", i, string(s[i]))
			i++
		}
		fmt.Printf("\n====================================\n")

	s_low := strings.ToLower(s)
	ind := strings.LastIndex(s_low, "n")
	fmt.Printf("lenght of s: %d\n", len(s))
	fmt.Printf("last index of n: %d\n", ind)
	x := string(s[ind])
	//x := strings.ToLower(string(s[len(s)-1:]))
	fmt.Printf("endswithn (x: %s, n: %s) \n ", x, n)
	return x == n
}
*/
func endsWithN(s []string) bool {
	var n string = "n"
	//s_ := string(strings.Replace(s, " ", "", -1))
	/*
		fmt.Printf("\n====================================\n")
		fmt.Printf("lenght of s: %d", len(s))
		for i := 0; i < len(s)-2; {
			fmt.Printf("(i: %d, char: %s)\n", i, string(s[i]))
			i++
		}
		fmt.Printf("\n====================================\n")
	*/
	fmt.Printf("\n====================================\nendsWithN\n")
	//s_low := strings.ToLower(s[len(s)-1])
	//ind := strings.LastIndex(s_low, "n")
	fmt.Printf("lenght of s: %d\n", len(s))
	//fmt.Printf("last index of n: %d\n", ind)
	x := strings.ToLower(string(s[len(s)-1]))
	//x := strings.ToLower(string(s[len(s)-1:]))
	fmt.Printf("endswithn (x: %s, n: %s) \n ", x, n)
	fmt.Println(x == n)
	fmt.Printf("\n====================================\n")
	return x == n
}

/*
func findLastCharacter(s string) string {
	for i := 0; i < len(s); i++ {

	}
}
*/
/*
func hasA(s string) bool {
	var a string = "a"
	x := strings.ToLower(s)
	fmt.Printf("hasA (%s %s)\n", string(x), a)
	return strings.Contains(x, a)
}
*/
func hasA(s []string) bool {
	var a string = "a"
	fmt.Printf("\n====================================\n")
	x := strings.ToLower(strings.Join(s, ""))
	fmt.Printf("hasA (%s %s)\n", string(x), a)
	fmt.Printf("\n====================================\n")
	return strings.Contains(x, a)
}
func removeSpaces(s string) ([]string, int) {
	N := len(s)
	fmt.Println(N)
	var s_ []string
	var qtd int
	for i := 0; i < N; i++ {

		if string(s[i]) != " " && string(s[i]) != "\n" {
			fmt.Printf("(%s %d)\n", string(s[i]), i)
			s_ = append(s_, string(s[i]))
		} else {
			qtd++
		}
	}
	fmt.Printf("\nqtd:%d\n", qtd)
	return s_, qtd
}

func main() {
	var s_input string

	fmt.Printf("Please enter a string and press ENTER.\n")

	//num, err := fmt.Scan(&s)
	//var num string

	in := bufio.NewReader(os.Stdin)
	s_input, err := in.ReadString('\n')

	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//s := scanner.Text()

	if err == nil {

		fmt.Printf("SCANNED s: %s \n", s_input)

		//s_ := string(strings.ReplaceAll(string(s), " ", ""))
		//s_, qtd := removeSpaces(strings.ReplaceAll(s_input, "\n", ""))
		s, qtd := removeSpaces(strings.ReplaceAll(s_input, "\n", ""))
		//s_ := string(s_nospace[:len(s_nospace)-qtd])

		//fmt.Printf("replaced s: %s \n", s_[:len(s_)-qtd])

		//s := s_[:len(s_)-qtd]
		fmt.Printf("replaced s: %s \n", s[:len(s)-qtd])

		if startsWithI(s) && endsWithN(s) && hasA(s) {
			fmt.Printf("Found!!\n")
		} else {

			fmt.Printf("\n====================================\n")
			fmt.Printf("\n====================================\n")
			fmt.Printf("\n====================================\n")
			fmt.Printf("Not Found!!\n")
			fmt.Println("startsWithI: ", startsWithI(s))
			fmt.Println("endsWithN(s): ", endsWithN(s))
			fmt.Println("hasA(s): ", hasA(s))
		}

	} else {
		//fmt.Printf("num: %d\n", num)
		fmt.Printf("error: %s\n", err)
	}

}
