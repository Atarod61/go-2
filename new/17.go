package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var A byte = 65 // مقدار ASCII کاراکتر 'A'

	fmt.Printf("value: %c, size: %d byte\n", A, unsafe.Sizeof(A))
}
