package main

import "fmt"

func iniSaya() (depan string, tengah string, belakang string) {
	depan = "Muhamad"
	tengah = "Wildan"
	belakang = "Faz"
	return depan, tengah, belakang
	//bisa juga tanpa menyebut depan, tengah, belakang setelah return
}

func main() {
	namaDepan, namaTengah, namaBelakang := iniSaya()
	fmt.Println(namaDepan, namaTengah, namaBelakang)
}
