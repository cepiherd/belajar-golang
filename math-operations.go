package main

import "fmt"

func main() {
	var rest = 10 + 10
	fmt.Println(rest)

	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	// augmented
	a += a
	fmt.Println(a) //output 20

}
