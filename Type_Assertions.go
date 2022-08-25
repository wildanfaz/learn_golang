package main

import "fmt"

func ganti() interface{} {
	return "Halo"
}

func main() {
	var aku interface{} = ganti()
	var ubah string = aku.(string)
	fmt.Println(ubah)

	switch cekType := aku.(type) {
	case string:
		fmt.Println("Ini adalah string", cekType)
	case int:
		fmt.Println("Ini adalah int", cekType)
	default:
		fmt.Println("Mbuh")
	}
}
