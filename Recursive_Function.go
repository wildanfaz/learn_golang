package main

import "fmt"

func faktor(apple int) int {
	hasil := 1
	for a := apple; a > 0; a-- {
		hasil *= a
	}
	return hasil
}

func rekursip(km int) int {
	if km == 1 {
		return 1
	} else {
		return km * rekursip(km-1)
	}
}

func main() {
	fmt.Println(faktor(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(rekursip(5))
}
