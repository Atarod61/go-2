package main

import "fmt"

func main() {

	orginalNumbers := []int{2, 3, 5, 11, 13, 1, 12, 43, 65, 666, 192}
	// 0, 1, 2, 3,  4, 5,  6,  7,  8,  9,   10

	//COPY :
	numbers_Of_Orginal_Numbers := []int{0, 1, 2, 3}
	copy(orginalNumbers[3:], numbers_Of_Orginal_Numbers[0:3])

	fmt.Println(orginalNumbers)
	fmt.Println(numbers_Of_Orginal_Numbers)

	//Change a one Index in Slices
	//fmt.Println(numbers)
	//numbers[0] = 6266
	//fmt.Println(numbers)

	//fmt.Println(len(numbers[3]))

	//Reverse Indexing
	//fmt.Println("Last number:", numbers[len(numbers)-2])

	//Length 1, 2, 3, 4, 5, 6
	//fmt.Println(len(numbers))

	//Indexing
	//fmt.Println(numbers[2])
}
