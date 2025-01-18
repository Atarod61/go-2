package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var A byte = "A"
	fmt.Printf("value: %c, size: %d byte\n", A, unsafe.Sizeof(A))
}
