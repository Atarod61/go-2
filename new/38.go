package main

import "fmt"

func main() {
	mysum := sum(2, 2)
	fmt.Println(mysum)

}
func sum(a int, b int) int {
	return a + b
}
func Minus(a int, b int) int {
	return a - b
}
func Divided(a int, b int) float32 {
	return float32(a / b)

}
func Multiplied(a int, b int) int {
	return a * b
}
