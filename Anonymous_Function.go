package main

import "fmt"

type diblok func(string) bool

func blockk(nama string, blck diblok) {
	if blck(nama) {
		fmt.Println("Anda Diblokir")
	} else {
		fmt.Println("Halo", nama)
	}
}

func kenapaDiBlok(nama string) bool {
	return nama == "asu"
}

func main() {
	blokir := func(nama string) bool {
		return nama == "goblog"
	}
	blockk("Bang", kenapaDiBlok)
	blockk("asu", kenapaDiBlok)
	blockk("goblog", blokir)
	blockk("titid", func(s string) bool {
		return s == "titid"
	})
}
