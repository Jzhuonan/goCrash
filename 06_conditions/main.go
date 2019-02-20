package main

import "fmt"

func main() {

	x := 6
	y := 6

	if x > y {
		fmt.Printf("%d is more than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	color := "yellow"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	case "green":
		fmt.Println("Color is green")
	default:
		fmt.Println("This is yoru favorite color")
	}
}