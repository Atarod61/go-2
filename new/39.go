package main

import "fmt"

func main() {
	r, b := divided(10, 3)
	fmt.Println(r, b)

}
func Divided(a int, b int) (int, int) {
	result := a / b
	baghimande := a % b
	return result, baghimande
}
