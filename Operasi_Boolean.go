package main

import "fmt"

func main() {
	var nilai = 90
	var absen = 90
	var ujian = 70
	var hasilA = nilai >= 80
	var hasilB = absen >= 80
	var hasilC = ujian >= 80
	fmt.Println(hasilA && hasilB)
	fmt.Println(hasilA || hasilC)
	fmt.Println(nilai >= 80 && absen >= 80)
}
