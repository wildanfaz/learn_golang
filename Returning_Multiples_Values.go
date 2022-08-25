package main

import "fmt"

func dualSword() (string, string) {
	return "Red Sword", "Blue Sword"
}

func namaneNyong() (string, string, string) {
	return "Muhamad", "Wildan", "Faz"
}

func main() {
	sword1, sword2 := dualSword()
	fmt.Println(sword1)
	fmt.Println(sword2)

	_, tengah, _ := namaneNyong()
	fmt.Println(tengah)
}
