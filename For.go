package main

import "fmt"

func main() {
	nomer := 1
	for nomer <= 10 {
		fmt.Println("Perulangan Ke", nomer)
		nomer++
	}

	for angka := 1; angka <= 3; angka++ {
		fmt.Println("Astaghfirullah")
	}

	sliceee := []string{"Muhamad", "Wildan", "Faz"}
	for i := 0; i < len(sliceee); i++ {
		fmt.Println(sliceee[i])
	}

	for _, nama := range sliceee {
		//index
		//fmt.Println("Index", index, "Nama", nama)
		fmt.Println(nama)
	}

	aku := make(map[string]string)
	aku["Nama"] = "Muhamad Wildan Faz"
	aku["Tittle"] = "Programmer"

	for key, value := range aku {
		fmt.Println(key, "=", value)
	}
}
