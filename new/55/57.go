package main

import "fmt"

type MyError struct {
	code    int
	message string
}

func (e MyError) Error() string {
	fmt.Sprint("Error: %v\")
}

func main() {

}
