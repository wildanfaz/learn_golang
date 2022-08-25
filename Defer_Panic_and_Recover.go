package main

import "fmt"

func oke() {
	fmt.Println("Wa'alaikumussalam")
}

func mantap() {
	//dengan recover menangkap panic tidak menghentikan program berikutnya
	pesan := recover()
	if pesan != nil {
		fmt.Println("Pesan :", pesan)
	}
	//defer tetep jalan walaupun func sedurunge error
	defer oke()
	fmt.Println("Assalamu'alaikum")
}

func panik(ngebug bool) {
	defer mantap()
	if ngebug {
		panic("Lah ko ngebug")
	} else {
		fmt.Println("Alhamdulillah Lancar Bang")
	}
}

func main() {
	mantap()
	//dong error aja panikan bang sing tenang tapi dokoni panic() karo dokoni recover()
	panik(false)
	fmt.Println("Mantap")
}
