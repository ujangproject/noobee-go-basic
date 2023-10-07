package main

import "fmt"

func main() {
	num := 5

	if num%2 == 0 {
		fmt.Printf("%d Merupakan bilangan genap.\n", num)
	} else {
		fmt.Printf("%d Merupakan bilangan ganjil.\n", num)
	}

}
