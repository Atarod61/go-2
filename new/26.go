package main

import "fmt"

func main() {
	//for i := 1; i <= 10; i++ {
	//fmt.Println(i)
	//}
	//i := 1
	//for i <= 10 {
	//fmt.Println(i)
	//i++
	//}
	//i:= 1
	//for{
	//fmt.Println(i)
	//i++
	//if i>10{
	//break
	//}
	//}
	//i := 1
	//for {
	//fmt.Println(i)
	//	i++
	//if i > 10 {
	//break
	//}

	//}
	//num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//for i, nums := range num {
	//	fmt.Println("index", i, "var in arrey", nums)
	//}
	mymap := map[string]string{"a": "apple", "b": "banana", "c": "chery"}
	for k, v := range mymap {
		fmt.Println("key:", k, "value:", v)
	}

}
