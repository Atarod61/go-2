package main

import "fmt"

func main() {
	var number int = 100
	var mypointer *int = &number
	fmt.Println(&number)
	fmt.Println(mypointer)
}
