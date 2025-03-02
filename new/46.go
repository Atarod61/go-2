package main

import "fmt"

type car struct {
	Brand string
	model string
	Year  int
}

func (c car) Display() string {
	return fmt.Sprintf("Brand:%s, model:%s, Year:%d", c.Brand, c.model, c.Year)

}

func main() {
	mycar := car{Brand: "TOyoTa", model: "corolla", Year: 2020}
	info := mycar.Display()
	fmt.Println(info)

}
