package main

import "fmt"

func main() {
	myarry := []int{12, 13, 14, 15}
	myslice := myarry[:]
	//println(myslice)
	fmt.Printf("%v\n", myslice)
}
