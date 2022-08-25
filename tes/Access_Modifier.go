package tes

import "fmt"

func Lalo(nama string) {
	fmt.Println("Halo", nama)
}

//gunakan awalan huruf besar pada function atau variabel jika akan diekspor ke package lain
//jika awalan huruf kecil tidak bisa diekspor
func halo(nama string) {
	fmt.Println("Halo", nama)
}

var Nama = "Muhamad Wildan Faz"

func Oke(nama string) {
	fmt.Println("Halo")
}
