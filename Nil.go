package main

import "fmt"

func gaweMap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"Nama": nama,
		}
	}
}

func main() {
	var a map[string]string = nil
	aa := map[string]string(nil)
	fmt.Println(a)
	fmt.Println(aa)

	aku := gaweMap("Wildan")
	fmt.Println(aku)

	var nama map[string]string
	if nama["Nama"] == "" {
		fmt.Println("Namane Kosong")
	} else {
		fmt.Println(nama)
	}

	cek := map[string]string(gaweMap(""))
	if cek == nil {
		fmt.Println("Namane Durung Diisi")
	} else {
		fmt.Println(cek)
	}
}
