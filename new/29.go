package main

import "fmt"

func main() {

	myarray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myarray)
	//println(&myarray)
	myslice := myarray[:]
	//println(myslice)
	fmt.Println(myslice)
	myarray[1] = 100
	myarray[2] = 200
	fmt.Println(myarray)
	fmt.Println(myslice)

}
