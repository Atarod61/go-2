package main

import "fmt"

func main() {

	myarray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myarray)
	println(&myarray)
	myslice := myarray[:]
	println(myslice)

}
