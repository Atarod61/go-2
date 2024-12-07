package main

import "fmt"

func main() {

	numbers := []int{2, 3, 5, 11, 13, 1}
	// 0, 1, 2, 3,  4, 5
	fmt.Println(numbers)
	numbers[0] = 6266
	fmt.Println(numbers)

	//fmt.Println(len(numbers[3]))

	//Reverse Indexing
	//fmt.Println("Last number:", numbers[len(numbers)-2])

	//Length 1, 2, 3, 4, 5, 6
	//fmt.Println(len(numbers))

	//Indexing
	//fmt.Println(numbers[2])
}
