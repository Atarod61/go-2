package main

import "fmt"

func main() {
	r, b, m := Divide(10, 3)
	fmt.Println(r, b, m)

}
func Divide(a int, b int) (int, int, string) {
	result := a / b
	baghimande := a % b
	var messsege string
	if baghimande != 0 {
		messsege = "baghimande darad bakhsh pazir nist"
	} else {
		messsege = "baghimande nadarad bakhsh pazir hast!"
	}

	return result, baghimande, messsege

}
