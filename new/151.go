package main

import "fmt"

func main() {
	var number int = 100
	var mypointer *int = &number
	fmt.Println(number)
	sum(mypointer)
	fmt.Println(number)

}
func sum(pointer *int) {
	*pointer = *pointer + 200
}
