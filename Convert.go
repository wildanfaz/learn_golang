package main

import "fmt"

func main() {
	var nilai1 int32 = 25000
	var nilai2 int64 = int64(nilai1)
	var nilai3 int8 = int8(nilai1)
	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)

	var nama = "Wildan"
	var w byte = nama[0]
	var setering string = string(w)
	fmt.Println(nama)
	fmt.Println(w)
	fmt.Println(setering)
}
