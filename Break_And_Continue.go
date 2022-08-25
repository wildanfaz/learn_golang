package main

import "fmt"

func main() {
	for tes := 1; tes <= 5; tes++ {
		fmt.Println("Perulangan", tes)
		if tes == 3 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Nomer Ganjil", i)
	}
}
