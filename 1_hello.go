package main

import "fmt" //format

func main() {
	var x string = "Hello"
	endereco := &x
	dados := *endereco

	var ip *int
	// fmt.Println(x)
	fmt.Printf("x: %s\n", x)
	fmt.Printf("endereco: %s\n", endereco)
	fmt.Printf("dados: %s\n", dados)
	fmt.Printf("ip: %s\n", ip)
}
