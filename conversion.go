package main

import "fmt"

func main() {
	var nilai32 int32 = 100000         //:=> output nya 100000
	var nilai64 int64 = int64(nilai32) //=> outpunya 100000
	var nilai8 int8 = int8(nilai64)    //outputnya -96 karena int8 tidak bisa memuat 100000

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "cepi"
	var e byte = name[0]           //read int8
	var eString string = string(e) //conversion dari int8 menjadi string

	fmt.Println(name)
	fmt.Println(eString)
}
