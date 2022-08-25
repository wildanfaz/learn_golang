package main

import "fmt"

type saya struct {
	nama, hobi string
	umur       int
}

func main() {
	var data saya
	data.nama = "Muhamad Wildan Faz"
	data.hobi = "Menikmati Waktu Bersamanya"
	data.umur = 18
	fmt.Println(data.nama)
	fmt.Println(data.hobi)
	fmt.Println(data.umur)
	fmt.Println(data)

	//var iya saya = saya{
	//	nama: "Muhamad Wildan Faz",
	//	hobi: "Menikmati Waktu Bersamanya",
	//	umur: 18,
	//}
	//fmt.Println(iya)

	adalah := saya{
		nama: "Muhamad Wildan Faz",
		hobi: "Menikmati Waktu Bersamanya",
		umur: 18,
	}
	fmt.Println(adalah)

	kenapa := saya{"Muhamad Wildan Faz", "Menikmati Waktu Bersamanya", 18}
	fmt.Println(kenapa)
}
