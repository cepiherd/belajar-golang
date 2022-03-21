package main

import "fmt"

func main() {
	var name1 = "cepi"
	var name2 = "cepi"

	var result bool = name1 == name2
	fmt.Println(result)

	var v1 = 100
	var v2 = 200

	fmt.Println(v1 > v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 == v2)
	fmt.Println(v1 != v2)
}
