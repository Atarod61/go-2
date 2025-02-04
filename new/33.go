package main

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice1 = append(slice1, slice2...)
}
