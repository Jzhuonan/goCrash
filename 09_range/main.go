package main

import "fmt"

func main() {

	randnums := []int{23, 34, 87, 2, 5, 48}

	// for i, v := range randnums {
	// 	fmt.Println(i, v)
	// }
	sum := 0
	for _, v := range randnums {
		sum += v
	}
	// fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string]string{"1": "thedog@gmail.com", "2": "jack@gmail.com", "3": "tom@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s => %s\n", k, v)
	}

	for _, v := range emails {
		fmt.Println(v)
	}
}