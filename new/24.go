package main

import "fmt"

func main() {
	//age1:= 17
	//age2 := 17
	//if age2 >= 18 {
	//println("yes! you can vote")
	//} else if age2 == 17 {
	//println("no! next year you can vote")

	//} else {
	//println("no!you can vote")
	//}
	//num1 := 10
	//num2 := 20
	//if num1 <= num2 {
	//fmt.Println("num1 is less than or equal to num2")
	//}
	//if num1 > num2 {
	//fmt.Println("num1 is greater than num2")
	//}
	//if num1 == num2 {
	//fmt.Println("num1 is equal to num2")
	//}
	//if num1 <= num2 && num1 == 11 {
	//fmt.Println("num1 is less than or equal to num2")

	//}
	//if num1 >= num2 || num1 == 10 {
	//fmt.Println("num1 is greater than num2")
	x := 2
	switch x {
	case 1:
		fmt.Println("x eqals 1")
	case 2:
		fmt.Println("x eqals 2")
	case 3:
		fmt.Println("x eqals 3")
	case 4:
		fmt.Println("x eqals 4")
	case 5:
		fmt.Println("x eqals 5")
	default:
		fmt.Println("x does not eqals 1, 2, 3, 4, 5")
	}

}
