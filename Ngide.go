package main

import "fmt"

func ngide(a int, b int) (int, int) {
	return 123, 123
}

func ooo(a int) error {
	return nil
}
func main() {
	a, b := ngide(1, 2)
	fmt.Println(a, b)

	o := ooo(123)
	if o != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("Aman")
	}
}
