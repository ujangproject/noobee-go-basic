package main

import "fmt"

func main() {
	gender := "male"
	age := 19

	if gender == "female" || age >= 21 {
		fmt.Println("Pelamar Dapat Dipekerjakan")
	} else {
		fmt.Println("Pelamar Tidak Dapat Dipekerjakan")

	}

}
