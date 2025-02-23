package main

import (
	"fmt"
	"os"
)

func main() {

	//defer fmt.Println("this is first")
	//fmt.Println("this is second")
	//defer fmt.Println("this is third")
	file, err := os.Open("test.tex")
	if err != nil {
		fmt.Println("my err:", err)
		return
	}
	fmt.Println("file opened:", file)
}
