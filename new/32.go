package main

import "fmt"

func main() {
	slices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//newslice := make([]int, len(slices))
	//fmt.Printf("this is slice %d\n", newslice)
	//n := copy(newslice, slices)
	//println(n)
	//fmt.Printf("this is a slice %d", newslice)
	fmt.Printf("this is a slice %d", slices)
	myfunc(slices)
	fmt.Printf("this is a slice after change in function %d", slices)

}
func myfunc(myslice []int) {
	slice[0] = 10000
	slice[2] = 2323233
}
