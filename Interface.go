package main

import "fmt"

type okei interface {
	ambilNama() string
}

func chat(iya okei) {
	fmt.Println("Insya Allah", iya.ambilNama())
}

func (nama hebat) ambilNama() string {
	return nama.aku
}

type hebat struct {
	aku string
}

func (nama hobi) ambilNama() string {
	return nama.suka
}

type hobi struct {
	suka string
}

func main() {
	data := hebat{aku: "Muhamad Wildan Faz"}
	fmt.Println(data.ambilNama())
	chat(data)

	kesukaan := hobi{suka: "Dapat Meraih Cita-Cita"}
	chat(kesukaan)
}
