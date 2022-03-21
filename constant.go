package main

import "fmt"

func main() {
	// constant jika data sudah di deklarasikan maka tidak dapat di ubah lagi
	const (
		firstname = "cepi"
		lastname  = "herdiansyah"
		value     = 1000
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)
}
