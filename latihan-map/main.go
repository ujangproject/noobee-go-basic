package main

import "fmt"

func main() {
	phoneBook := make(map[string]string)

	phoneBook["tio"] = "087810990165"
	phoneBook["sigit"] = "08991980409"

	fmt.Println("Semua Kontak", phoneBook)
	fmt.Println("Nomor Telepon Tio", phoneBook["tio"])

	phoneBook["aldi"] = "081210990687"

	fmt.Println("Setelah menambahkan kontak Aldi", phoneBook)

	delete(phoneBook, "sigit")

	fmt.Println("Setelah hapus kontak sigit", phoneBook)

	fmt.Println("Data Kontak")

	for key, value := range phoneBook {
		fmt.Printf("%s:%s\n", key, value)
	}
}
