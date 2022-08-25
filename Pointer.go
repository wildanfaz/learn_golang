package main

import "fmt"

type lokasi struct {
	rt, rw int
	desa   string
}

func main() {
	//a := lokasi{04, 03, "Pepedan"}
	//b := a

	//a := lokasi{04, 03, "Grogol"}
	//b := &a

	var a lokasi = lokasi{04, 03, "Pepedan"}
	var b *lokasi = &a
	var c *lokasi = &a
	
	//var d *lokasi = &lokasi{04, 03, "Pepedan"}
	var d *lokasi = new(lokasi)
	d.desa = "Pepedan"
	fmt.Println(d)

	b.desa = "Grogol"
	//pointer hanya bisa merubah field

	fmt.Println(a)
	fmt.Println(b)

	//b = &lokasi{1, 2, "Mbuh"}
	//operator & membuat memori baru
	//a tidak berubah
	//fmt.Println(a)
	//fmt.Println(b)

	*b = lokasi{04, 03, "Pepedan"}
	//a ikut berubah dengan menggunakan operator *
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
