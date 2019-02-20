package main

import "fmt"

func main() {

	a := 3.1
	b := &a

	fmt.Println(*b)
	// fmt.Println(*&a)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	*b = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}