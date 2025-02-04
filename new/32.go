package main

import (
	"fmt"
)

func main() {
	slices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//newslice := make([]int, len(slices))
	//fmt.Printf("this is slice %d\n", newslice)
	//n := copy(newslice, slices)
	//println(n)
	//fmt.Printf("this is a slice %d", newslice)
	fmt.Printf("this is a slice %d\n", slices)
	myfunc(slices)
	fmt.Printf("\n this is a slice after change in function %d", slices)

}
func myfunc(myslice []int) {
	newslice := make([]int, len(myslice))
	copy(newslice, myslice)
	newslice[0] = 10000
	newslice[2] = 232323
}
