package main

import "fmt"

func main() {
	//Blok Selanjute Kudu Ngarepe Kurawal {}
	nama := "Muhamad Wildan Faz"
	if nama != "Muhamad Wildan Faz" {
		fmt.Println("Namane Salah")
	} else if nama == "" {
		fmt.Println("Namane Durung Diisi")
	} else {
		fmt.Println("Namane Wis Bener")
	}

	//panjang := len(nama)
	if panjang := len(nama); panjang > 10 {
		fmt.Println("Nama Kedawanen")
	} else if panjang < 3 {
		fmt.Println("Namane Kependeken")
	} else {
		fmt.Println("Namane Wis Bener")
	}
}
