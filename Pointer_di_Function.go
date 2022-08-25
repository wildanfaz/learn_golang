package main

import "fmt"

type dataku struct {
	nama string
	umur int
}

func person(a *dataku) {
	a.umur = 18
}

func main() {
	abc := dataku{"Muhamad Wildan Faz", 0}
	person(&abc)
	fmt.Println(abc)
}
