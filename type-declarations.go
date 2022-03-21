package main

import "fmt"

func main() {
	type NoKTP string //merubah type data yang sudah ada menjadi yang kita inginkan agar lebih mudah di pahami.

	var NoKTPCepi NoKTP = "918293821933"
	fmt.Println(NoKTPCepi)
}
