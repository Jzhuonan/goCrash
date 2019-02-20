package main

import "fmt"

// var name = "theDog"


func main() {
	// Main Types
	// string
	// bool
	// int
	// int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// var name = "theDog"
	// var age int64 = 24
	// var isCool = true
	// isCool = false

	// name := "theDog"
	// email := "jzhuonan@gmail.com"

	name, email := "theDog", "jzhuonan@gmail.com"	
	// size := 1.3

	fmt.Printf("%T - %T\n", name, email)
	fmt.Println(name, email)

}

