package main

import "fmt"

func main() {
	coba := map[string]string{
		"Nama": "Muhamad Wildan Faz",
		"Umur": "18",
	}
	fmt.Println(coba)
	fmt.Println(coba["Nama"])
	coba["Tittle"] = "Programmer"
	coba["Hobi"] = "Rebahan"
	fmt.Println(coba)

	baru := make(map[string]string)
	baru["Nama"] = "Muhamad Wildan Faz"
	fmt.Println(baru)

	delete(coba, "Hobi")
	fmt.Println(coba)
	fmt.Println(len(coba))
}
