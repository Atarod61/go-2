package main

import (
	"fmt"
)

func main() {

	//defer fmt.Println("this is first")
	//fmt.Println("this is second")
	//defer fmt.Println("this is third")
	//file, err := os.Open("test.tex")
	//	defer file.Close()
	//if err != nil {
	//fmt.Println("my err:", err)
	//return
	//}
	//fmt.Println("file opened:", file)
	x := 10
	defer fmt.Println("this is whit defer", x)

	x = 20
	fmt.Println("this is whitout defer", x)

}
