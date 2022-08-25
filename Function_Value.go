package main

import "fmt"

func getHalo(nama string) string {
	return "Halo " + nama
}

func main() {
	//Aja dalai () ngko error
	halo := getHalo
	hasil := halo("Wildan")
	fmt.Println(hasil)
}
