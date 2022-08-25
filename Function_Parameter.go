package main

import "fmt"

func namaSaya(depan string, tengah string, belakang string) {
	fmt.Println("Nama Saya", depan, tengah, belakang)
}

func main() {
	tengah := "Wildan"
	namaSaya("Muhamad", tengah, "Faz")
}
