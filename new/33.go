package main

import (
	"fmt"
)

func main() {
	//slice1 := []int{1, 2, 3}
	//slice2 := []int{4, 5, 6}
	//slice1 = append(slice1, slice2...)
	//fmt.Println(slice1)
	numbers := []int{1, 2, 3}
	part1 := numbers[:2]
	part2afterappend := append([]int{10}, numbers[2:]...)
	atlast := append(part1, part2afterappend...)
	fmt.Println(atlast)
}
