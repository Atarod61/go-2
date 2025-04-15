package main

import (
	"errors"
	"fmt"
)

//"encoding/json"
//"fmt"

func main() {
	result, error := createError(true)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(result)
	//result := struct {
	//name string
	//age  int
	//}{
	//name: "alireza",
	//age:  27,
	//}
	//studenjson, err := json.Marshal(result)
	//if err != nil {
	//fmt.Println("Error marshaling", err)
	//}
	//fmt.Println(string(studenjson))
}
func createError(er bool) (string, error) {
	if er {
		return "", errors.New("error creating")
	}
	return "you dont like create error", nil
}
