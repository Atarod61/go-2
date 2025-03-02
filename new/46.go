package main

import "fmt"

type car struct {
	Brand string
	model string
	Year  int
}

func (c car) Display() string {
	return fmt.Sprintf("Brand:%s, model:%s, year:%d", c.Brand, c.model, c.Year)

}

func main() {
	car:={Brand:"TOyoTa", Model:"corolla", year:2020}
	info:= car.Display(
		fmt.Println(info)
	)

}
