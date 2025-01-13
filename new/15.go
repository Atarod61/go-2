package main

import "fmt"

func main() {
	var number int = 100
	var mypointer *int = &number
	fmt.Println("This is number address in memory:", &number)
	fmt.Println("This is mypointer address in memory:", mypointer)
	fmt.Println(".........")
	fmt.Println("This is number value:", number)
	fmt.Println("This is mypointer value:", mypointer)
}
