package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "Muhamad"
	nama[1] = "Wildan"
	nama[2] = "Faz"
	fmt.Println(nama)
	fmt.Println(nama[1])

	var nomer = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(nomer)
	fmt.Println(nomer[1])

	fmt.Println(len(nama))
	nomer[2] = 12
	fmt.Println(nomer[2])
}
