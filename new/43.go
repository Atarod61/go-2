package main

import "fmt"

func main() {

	defer fmt.Println("this is first")
	fmt.Println("this is second")
	defer fmt.Println("this is third")
}
