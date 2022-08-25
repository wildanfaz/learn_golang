package main

import "fmt"

type aku struct {
	nama, hobi string
	umur       int
}

func (dataku aku) halo(nama string) {
	fmt.Println("Halo", nama, "Nama Saya", dataku.nama)
}

func (halo aku) okey() {
	fmt.Println("Halo", halo.nama)
}

func main() {
	adalah := aku{
		nama: "Muhamad Wildan Faz",
		hobi: "Menikmati Waktu Bersamanya",
		umur: 18,
	}
	fmt.Println(adalah)
	adalah.halo("Mba")
	adalah.okey()
}
