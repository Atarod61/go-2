package main

import "fmt"

type persons struct {
	name string
	Id   int
	age  int
}
type Employee struct {
	code      string
	TIMESTAMP int64
	persons
}

func main() {

	e := Employee{
		code:      "12",
		TIMESTAMP: 10,
		persons: persons{
			name: john,
			Id:   1,
			age:  25,
		},
	}
	M := Maneger{
		code:    "1234",
		company: "cocompany",
		persons: persons{
			name: "max",
			Id:   2,
			age:  45,
		},
	}
	fmt.Println(e)
	fmt.Println(M)
}
