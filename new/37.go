package main

import "fmt"

func main() {
	//num1 := 10
	//num2 := 10
	//result := num1 + num2
	//fmt.Println(result)
	result := sum(10, 10)
	fmt.Println(result)

}
func sum(number1 int, number2 int) int {
	return number1 + number2
}
