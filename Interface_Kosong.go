package main

import "fmt"

func bebas(nomer int) interface{} {
	if nomer == 1 {
		return "Mantap"
	} else if nomer == 2 {
		return true
	} else {
		return 123
	}
}

func main() {
	fmt.Println(bebas(1))
	save := bebas(2)
	fmt.Println(save)
	var simpan interface{} = bebas(3)
	fmt.Println(simpan)
}
