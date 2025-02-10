package main

import "fmt"

func main() {
	personmap := map[int]string{1234: "Ali emadi", 234567: "mohamad ranjbar"}
	fmt.Println(personmap[1234])
	fmt.Println(personmap[234567])

}
