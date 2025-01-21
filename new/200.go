package main

import "fmt"

var GLOBAL int = 30

func main() {
	fmt.Println(GLOBAL)
	test()
}
func test() {
	var num int = 30
	GLOBAL = GLOBAL + num
	fmt.Println(GLOBAL)
}
