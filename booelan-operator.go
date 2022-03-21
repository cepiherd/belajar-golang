package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian bool = ujian >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusUjian && lulusAbsensi

	fmt.Println(lulus)

	// cara cepat

	fmt.Println(ujian >= 80 && absensi >= 80)
}
