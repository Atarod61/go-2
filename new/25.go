package main

func main(){
	x:= 71
	witch x {
	case 1: fmt.Println("x eqals 1")
	case 2: fmt.Println("x eqals 2")
	case 3: fmt.Println("x eqals 3")
	case 4:	fmt.Println("x eqals 4")
	case 5: fmt.Println("x eqals 5")
	default: fmt.Println("x does not eqals 1, 2, 3, 4, 5")
	}
}
