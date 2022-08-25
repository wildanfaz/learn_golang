package main

import "fmt"

func haloMas(namaku string) string {
	return "Halo " + namaku
}

func jikaMemang(aku string) string {
	if aku == "" {
		return "Harus Cari Uang"
	} else {
		return "Semangat " + aku
	}
}

func main() {
	halo := haloMas("Wildan")
	fmt.Println(halo)
	semangat := jikaMemang("Wildan")
	fmt.Println(semangat)
	fmt.Println(jikaMemang(""))
}
