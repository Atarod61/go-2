package main

import "fmt"

func main() {
	//var arr[5]int
	//	var arr [5]int= [5]int{1, 2, 3, 4, 5}
	//arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr[2])
	//arr[2] = 55
	//fmt.Println(arr[2])
	//fmt.Println(len(arr))
	//for i, v := range arr {
	//	fmt.Println(i, v)
	//}
	//arr := [5]int{1, 2, 3, 4, 5}
	//for i := 0; i < len(arr); i++ {
	//fmt.Println(arr[i])
	//}
	//	1 2 3
	//	4 5 6
	//7 8 9
	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(arr[1][2])
	fmt.Println(arr[1][1])
}
