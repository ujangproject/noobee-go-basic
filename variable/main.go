package main

import "fmt"

var school string = "SMK DARUSSALAM PANONGAN"

func main() {

	// Membuat variable

	var firstName string = "Muhammad Tio"
	var lastName string = "Laksono"

	var nickName string
	nickName = "Tio"

	myAge := 26

	fmt.Println("Full Name: ", firstName, lastName)
	fmt.Println("Nick Name: ", nickName)
	fmt.Println("Umur: ", myAge)

	// Variable Diluar Fungsi

	fmt.Println("School: ", school)

	// Variable Tanpa Nilai Awal

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Deklarasi Multiple Variable

	var d, e, f, g = 1, 2, 3, "Hello"

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Deklarasi Multiple Variable Dalam Blok

	var (
		alamat string = "Citra Raya, Cikupa Tangerang - Banten 15710"
		status bool   = true
		height int    = 162
	)

	fmt.Println("Alamat: ", alamat)
	fmt.Println("Status Sekolah: ", status)
	fmt.Println("Tinggi Badan: ", height)

}
