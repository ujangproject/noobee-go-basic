package main

import "fmt"

func main() {
	fruits := map[string]int{
		"Apple":  10,
		"Banana": 15,
		"Orange": 8,
		"Grape":  20,
	}

	fmt.Println("Sebelum ditambah/hapus")
	fmt.Println(fruits)
	fmt.Println()

	fruits["Manggo"] = 12
	fruits["Strawberry"] = 7

	fmt.Println("Setelah menambahkan manggo dan strawberry")
	fmt.Println(fruits)
	fmt.Println()

	delete(fruits, "Orange")

	fmt.Println("Setelah menghapus buah Orange")
	fmt.Println(fruits)
	fmt.Println()

	fmt.Println("Jumlah apel yang tersedia: ", fruits["Apple"])
	fmt.Println()

	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	fmt.Println()

	fmt.Println("Banana: ", fruits["Banana"])
	fmt.Println("Grape: ", fruits["Grape"])
	fmt.Println("Manggo: ", fruits["Manggo"])
	fmt.Println("Strawberry: ", fruits["Strawberry"])
	fmt.Println("Apple: ", fruits["Apple"])

}
