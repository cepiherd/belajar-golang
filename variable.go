package main

import "fmt"

func main() {
	var name string           // variable harus dengan type data nya jika ingin menentukan
	name = "Cepi Herdiansyah" //isi variables yang sudah di inisialisasikan bisa di ubah ubah
	fmt.Println(name)

	var sohibName = "abdullah"
	fmt.Println(sohibName)

	country := "indonesia"
	fmt.Println(country)

	country = "india"
	fmt.Println(country)

	// variable multiple lebih readable : recomended

	var (
		firstname = "abdullah"
		lastname  = "yahya"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}
