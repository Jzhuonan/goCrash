package main

import (
	"fmt"
	// "math"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(int1, int2 int) int {
	return int1 + int2
}

func main() {

	fmt.Println(greeting("theDog"))
	fmt.Println(getSum(2, 3))
}