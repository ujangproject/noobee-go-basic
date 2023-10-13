package main

import (
	"fmt"
)

func main() {

	// Latihan 1

	var d, e, f int
	f = 1
	for d = 5; d >= f; d-- {

		for e = 1; e <= d; e++ {

			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()
	// Latihan 2

	var a, b, c int
	c = 5
	for a = 1; a <= c; a++ {

		for b = 1; b <= a; b++ {

			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()
	// Latihan 3

	// var latihan_3 string = ""

	for i := 1; i <= 50; i++ {
		dibagi_3 := i % 3
		dibagi_5 := i % 5

		if dibagi_3 == 0 && dibagi_5 == 0 {
			fmt.Printf("NooBee, ")

		} else if dibagi_3 == 0 {
			fmt.Print("Noo, ")

		} else if dibagi_5 == 0 {
			fmt.Print("Bee, ")

		} else {
			fmt.Print(i, ", ")
		}
	}
}
