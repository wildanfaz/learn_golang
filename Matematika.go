package main

import "fmt"

func main() {
	var plus = 5 + 5
	fmt.Println(plus)
	var a = 5
	a += 5
	fmt.Println(a)
	var b = 9
	b++
	fmt.Println(b)
	var c = 11
	c--
	fmt.Println(c)
	var d = -10
	fmt.Println(d)
	var e = +10
	fmt.Println(e)
	fmt.Println(d + e)
}
