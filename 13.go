package main

import "fmt"

func main() {

	numbers := []int{2, 3, 5, 11, 13, 1}
	// 0, 1, 2, 3,  4, 5
	fmt.Println("Last number:", numbers[len(numbers)])

	//Length 1, 2, 3, 4, 5, 6
	//fmt.Println(len(numbers))

	//Indexing
	//fmt.Println(numbers[2])
}
