package main

import "fmt"

func main() {

	arr1 := [3]int{9, 7, 6}
	arr2 := [...]int{9, 7, 6, 4, 5, 3, 2, 4}
	arr3 := [3]int{9, 3, 5}

	fmt.Println("Length of the array 1 is:", len(arr1))
	fmt.Println("Length of the array 2 is:", len(arr2))
	fmt.Println("Length of the array 3 is:", len(arr3))

}
