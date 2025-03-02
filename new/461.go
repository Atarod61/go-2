package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func CaculateArea(width, height float64) float64 {
	return width * height
}

func (r Rectangle) Displayinfo() string {
	return fmt.Sprintf("Width:%.2f, Height: %.2f, Area: %.2f", r.width, r.height, CaculateArea(r.width, r.height))
}

func main() {
	r := Rectangle{width: 5, height: 3}
	result := r.Displayinfo()
	fmt.Println(result)
}
