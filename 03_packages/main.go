package main

import (
	"fmt"
	"math"

	"github.com/jzhuonan/go_crash_course/03_packages/strutil"
)

func main() {

	name := "theDog"
	age := 25
	sqrtAge := math.Sqrt(float64(age))

	fmt.Println(name, age, sqrtAge)
	fmt.Println(math.Floor(3.2))
	fmt.Println(math.Ceil(3.2))
	fmt.Println(strutil.Reverse("olleh"))
}