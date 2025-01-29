package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr2[0] = 100
	arr2[1] = 200
	arr2[2] = 300
	fmt.Printf("this is arr1 :%d , arr2:%d\n", arr1, arr2)
	print(&arr1, ",", &arr2)
}
