package main

import "fmt"

func main() {
	//var arr[5]int
	//	var arr [5]int= [5]int{1, 2, 3, 4, 5}
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2])
	arr[2] = 55
	fmt.Println(arr[2])

}
