package main

import "fmt"

type OrangTua struct {
	Nama string
	Umur int
}

type Siswa struct {
	OrangTua OrangTua
	Nama     string
	Umur     int
	Kelas    string
}

func main() {
	orangtua1 := OrangTua{
		Nama: "Budi",
		Umur: 40,
	}

	siswa1 := Siswa{
		OrangTua: orangtua1,
		Nama:     "Ali",
		Umur:     15,
		Kelas:    "9A",
	}

	orangtua2 := OrangTua{
		Nama: "Citra",
		Umur: 39,
	}

	siswa2 := Siswa{
		OrangTua: orangtua2,
		Nama:     "David",
		Umur:     14,
		Kelas:    "8B",
	}

	orangtua3 := OrangTua{
		Nama: "Eva",
		Umur: 35,
	}

	siswa3 := Siswa{
		OrangTua: orangtua3,
		Nama:     "Fina",
		Umur:     16,
		Kelas:    "10C",
	}

	fmt.Println("Informasi Siswa 1:")
	fmt.Printf("Nama: %s, Umur: %s, Kelas: %s\n", siswa1.Nama, siswa1.Umur, siswa1.Kelas)
	fmt.Printf("Orang Tua: %s, Umur: %d \n", siswa1.OrangTua.Nama, siswa1.OrangTua.Umur)
	fmt.Println()

	fmt.Println("Informasi Siswa 2:")
	fmt.Printf("Nama: %s, Umur: %s, Kelas: %s\n", siswa2.Nama, siswa2.Umur, siswa2.Kelas)
	fmt.Printf("Orang Tua: %s, Umur: %d \n", siswa2.OrangTua.Nama, siswa2.OrangTua.Umur)
	fmt.Println()

	fmt.Println("Informasi Siswa 3:")
	fmt.Printf("Nama: %s, Umur: %d, Kelas: %s\n", siswa3.Nama, siswa3.Umur, siswa3.Kelas)
	fmt.Printf("Orang Tua: %s, Umur: %d \n", siswa3.OrangTua.Nama, siswa3.OrangTua.Umur)
	fmt.Println()

}
