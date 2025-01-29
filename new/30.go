package main

import "fmt"

func main() {
	myarray := [5]int{1, 2, 3, 4, 5}
	myslice_fullarray := myarray[:]

	fmt.Println(myslice_fullarray)
	myslice_PartofArray := myarray[1:3]
	myslice_PartofArray2 := myarray[0:2] //1,2

	fmt.Println(myslice_PartofArray)
	fmt.Println(myslice_PartofArray2)

	myslice_PartofArray3 := myarray[2:]
	//fmt.Println(myslice_PartofArray3)
	//fmt.Println(len(myslice_PartofArray3))
	//fmt.Println(cap(myslice_PartofArray3))
	fmt.Println(myslice_PartofArray3[2])
}
