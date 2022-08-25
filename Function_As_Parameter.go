package main

import "fmt"

type a func(string) string

func harusSopan(namane string, sensor a) {
	hasil := sensor(namane)
	fmt.Println("Namane", hasil)
}

func sensorNama(nama string) string {
	if nama == "Asu" {
		return "..."
	} else {
		return nama
	}
}

func main() {
	aaa := sensorNama
	harusSopan("Asu", aaa)
	harusSopan("Wildan", aaa)
}
