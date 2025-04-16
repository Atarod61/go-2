package main

import "fmt"

type MyError struct {
	code    int
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %d: %s", e.code, e.message)
}
func ErroCretor(isError bool) (string, error) {
	if isError {
		return "", &MyError{code: 404, message: "not Found"}
	}
	return "we dont have Error", nil
}
func main() {
	result, error := ErroCretor(true) // تصحیح: استفاده از true و نام متغیر صحیح
	if error != nil {                 // تصحیح: استفاده از نام متغیر صحیح
		fmt.Println(error.Error())
	}
	fmt.Println(result)

}
