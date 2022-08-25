package main

import "fmt"

func main() {
	nama := "Wildan"
	switch nama {
	case "Wildan":
		fmt.Println("Namane Wis Bener")
	case "":
		fmt.Println("Namane Durung Diisi")
	default:
		fmt.Println("Halo")
	}

	switch namane := len(nama); namane > 10 {
	case true:
		fmt.Println("Namane Kedawanen")
	case false:
		fmt.Println("Namane Wis Bener")
	}

	panjang := len(nama)
	switch {
	case panjang > 10:
		fmt.Println("Nama Kedawan")
	case panjang < 1:
		fmt.Println("Nama Durung Diisi")
	default:
		fmt.Println("Namane Wis Bener")
	}
}
