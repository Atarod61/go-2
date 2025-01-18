package main

import "fmt"

func main() {

	var A byte = "A"
	fmt.Printf("value: %c, size: %d byte\n", A, unsafe.sizeof(A))
}
