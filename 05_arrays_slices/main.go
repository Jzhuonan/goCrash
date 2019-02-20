package main

import (
	"fmt"
)

func main() {
	// Array declaration and assignment
	// fruits := [3]string{"apple", "orange", "pear"}

	// Slice 
	fruits := []string{"apple", "orange", "pear", "grape"}


	fmt.Println(fruits)
	fmt.Println(fruits[1])

	fmt.Println(len(fruits))
	fmt.Println(fruits[1:3])
}