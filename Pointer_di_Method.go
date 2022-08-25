package main

import "fmt"

type okeBang struct {
	nama string
}

func (nyong *okeBang) okela() {
	nyong.nama = "Nama Saya " + nyong.nama
	fmt.Println(nyong)
}

func main() {
	a := okeBang{"Muhamad Wildan Faz"}
	a.okela()
	fmt.Println(a)
}
