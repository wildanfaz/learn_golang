package main

import (
	"errors"
	"fmt"
)

func bagi(nomer int, dibagi int) (int, error) {
	if dibagi == 0 {
		return 0, errors.New("Ora Bisa Dibagi 0")
	} else {
		hasil := nomer / dibagi
		return hasil, nil
	}
}

func main() {
	errorBang := errors.New("Error")
	fmt.Println(errorBang)

	nomer, bagi := bagi(10, 0)
	if bagi == nil {
		fmt.Println("Hasil", nomer)
	} else {
		fmt.Println("Error", bagi.Error())
	}
}
