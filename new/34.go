package main

import "fmt"

func main() {
	personmap := map[int]string{1234: "Ali emadi", 234567: "mohamad ranjbar"}
	fmt.Println(personmap[1234])
	fmt.Println(personmap[234567])
	fmt.Println(personmap[3456])

	personmap2 := map[int]string{}
	personmap2[124567] = "mostafa"
	personmap2[456789] = "hassan"

	fmt.Println(personmap2[124567])

}
