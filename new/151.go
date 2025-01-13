package main

import "fmt"

func main() {
	var number int = 100
	var mypointer *int = &number
	fmt.Println(number)
	sum(mypointe)
	fmt.Println(number)

}
func main(pointer *int) {
	*pointer = *pointer + 200
}
