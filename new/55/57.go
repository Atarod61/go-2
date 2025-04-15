package main

type MyError struct {
	code    int
	message string
}

func (e MyError) Error() string

func main() {

}
