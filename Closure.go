package main

import "fmt"

func main() {
	//global bisa diakses lokal, sedangkan lokal tidak bisa diakses global
	aku := "Wildan"
	hehe := func() {
		aku = "Faz"
		fmt.Println(aku)
		aku := "Muhamad Wildan Faz"
		fmt.Println(aku)
	}
	hehe()
}
