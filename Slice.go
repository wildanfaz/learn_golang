package main

import "fmt"

func main() {
	nomer := [...]int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	slice := nomer[2:8]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(nomer))

	//nomer[3] = 5
	//fmt.Println(slice)

	//slice[0] = 1000
	//fmt.Println(nomer)

	var slice2 = append(slice, 999)
	fmt.Println(slice2)
	fmt.Println(nomer)

	new := make([]string, 3, 6)
	new[0] = "Muhamad"
	new[1] = "Wildan"
	new[2] = "Faz"
	fmt.Println(new)

	var kopi = make([]string, len(new), cap(new))
	copy(kopi, new)
	fmt.Println(kopi)

	kieArray := [...]int{1, 2, 3}
	kieSlice := []int{1, 2, 3}
	fmt.Println(kieArray)
	fmt.Println(kieSlice)
}
