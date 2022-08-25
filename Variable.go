package main

import "fmt"

func main() {
	//var error jika tidak dipakai
	var nama string
	nama = "Muhamad Wildan Faz"
	fmt.Println(nama)

	var panggilan = "Wildan"
	var umur = 18
	fmt.Println(panggilan)
	fmt.Println(umur)

	negara := "Indonesia"
	fmt.Println(negara)

	var (
		namaDepan    = "Muhamad"
		namaTengah   = "Wildan"
		namaBelakang = "Faz"
	)
	fmt.Println(namaDepan, namaTengah, namaBelakang)
}
