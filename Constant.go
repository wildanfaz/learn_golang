package main

import "fmt"

func main() {
	//const tidak error jika tidak dipakai
	const name string = "Muhamad Wildan Faz"
	const umur = 18
	const (
		depan    string = "Muhamad"
		tengah          = "Wildan"
		belakang        = "Faz"
	)
	fmt.Println(name)
}
