package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func main() {
	book1 := Book{
		Title:  "Pemrograman Go",
		Author: "John Doe",
	}

	book2 := Book{
		Title:  "Algoritma dan Struktur Data",
		Author: "John Doe",
	}

	book3 := Book{
		Title:  "Machine Learning for Beginners",
		Author: "John Doe",
	}

	fmt.Println("Informasi tentang Book 1:")
	fmt.Println("Judul: ", book1.Title)
	fmt.Println("Penulis: ", book1.Author)
	fmt.Println()

	fmt.Println("Informasi tentang Book 2:")
	fmt.Println("Judul: ", book2.Title)
	fmt.Println("Penulis: ", book2.Author)
	fmt.Println()

	fmt.Println("Informasi tentang Book 3:")
	fmt.Println("Judul: ", book3.Title)
	fmt.Println("Penulis: ", book3.Author)
	fmt.Println()

}
