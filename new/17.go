package main

import (
	"fmt"
)

func main() {

	//var A byte = 65 // مقدار ASCII کاراکتر 'A'

	//fmt.Printf("value: %c, size: %d byte\n", A, unsafe.Sizeof(A))

	//var Farsi rune = 'پ' // استفاده از rune('پ') برای تبدیل کاراکتر فارسی به rune
	//fmt.Printf("value: %c, size: %d byte\n", Farsi, unsafe.Sizeof(Farsi))
	srt := "Hello"
	persiansrt := "سلام"
	fmt.Println(len(srt))
	fmt.Println(len(persiansrt))
}
