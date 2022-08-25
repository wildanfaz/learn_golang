package main

import "fmt"

func watashi(nomer ...int) int {
	total := 0
	for _, a := range nomer {
		total += a
	}
	return total
}

func main() {
	save := watashi(1, 2, 3, 4)
	fmt.Println(save)
	slice := []int{1, 2, 3, 4}
	fmt.Println(watashi(slice...))
}
