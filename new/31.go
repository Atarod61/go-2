package main

import "fmt"

func main() {
	//arr1 := [5]int{1, 2, 3, 4, 5}
	//arr2 := arr1
	//arr2[0] = 100
	//arr2[1] = 200
	//arr2[2] = 300
	//fmt.Printf("this is arr1 :%d , arr2:%d\n", arr1, arr2)
	//print(&arr1, ",", &arr2)
	myslice1 := []int{1, 2, 3, 4, 5}
	myslice2 := myslice1

	myslice2[0] = 100
	myslice2[2] = 300
	myslice2[3] = 600
	fmt.Printf("this is slice1 %d, slice2 %d\n", myslice1, myslice2)
	print(&myslice1, ",", &myslice2)
}
