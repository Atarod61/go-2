package main

import "fmt"

func main() {
	//x:= 3
	//witch x {
	//case 1: fmt.Println("sunday")
	//break
	//case 2: fmt.Println("Monday")
	//break
	//case 3: fmt.Println("Tuesday")
	//break
	//case 4:	fmt.Println("Wednesday")
	//break
	//case 5: fmt.Println("thursday")
	//break
	//case 5: fmt.Println("friday")
	//break
	//case 5: fmt.Println("saturday")
	//break
	//default: fmt.Println("we have worng")
	//}
	//Println("codes,,,,")
	x := 3
	switch x {
	case 1:
		fmt.Println("sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
		fallthrough
	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("thursday")

	case 6:
		fmt.Println("friday")

	case 7:
		fmt.Println("saturday")

	default:
		fmt.Println("we have worng")
	}
	fmt.Println("codes,,,,")
}
