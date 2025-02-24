package main

import "fmt"

type car struct {
	name       string
	createDate string
	company    string
}

func main() {
	//car1 := car{
	//name:       "Benz",
	//createDate: "2024",
	//company:    "Benz",
	//}
	//fmt.Println(car1)
	//car2 := &car{
	//name:       "Mercedes",
	//createDate: "2020",
	//company:    "Mercedes",
	//}
	//fmt.Println(car2)
	//carwithcomstracture := Newcar("paykan", "1898", "Irankhodro")
	//fmt.Println(carwithcomstracture)
	//car4 := car{name: "Pride"}
	//fmt.Println(car4)
	car5 := car{"Pride", "1999", "sipa"}
	fmt.Println(car5)

}

//func Newcar(name string, createDate string, companyname string) *car {
//return &car{name, createDate, companyname}
//}
