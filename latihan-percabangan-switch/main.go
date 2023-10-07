package main

import "fmt"

func main() {
	var huruf string = "a"

	switch huruf {
	case "a":
		fmt.Printf("Huruf %s adalah huruf vokal.\n", huruf)
	case "b":
		fmt.Printf("Huruf %s adalah huruf konsonan.\n", huruf)
	}

}
