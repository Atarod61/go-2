package main

import (
	"fmt"
)

func main() {
	slices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newslice := make([]int, len(slices))
	fmt.Printf("this is slice %d", newslice)

}
