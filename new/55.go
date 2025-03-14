package main

import "fmt"

func Print[T any](value T) T {
	fmt.Println(value)
	return value
}

func main() {
	Print(12)
	Print("hi my name is alireza")
	panic(12.5)

}
