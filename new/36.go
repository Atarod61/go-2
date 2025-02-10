package main

import "fmt"

func main() {
	//sa := map[int]string{1: "alireza", 2: "mohamad", 3: "maryam", 4: "yashar", 5: "mostafa", 6: "massod"}

	//fmt.Println(sa)
	//delete(sa, 4)
	//fmt.Println(sa)
	sa := map[int]string{1: "alireza", 2: "mohamad", 3: "maryam", 4: "yashar", 5: "mostafa", 6: "massod"}
	for key, value := range sa {
		fmt.Printf("key: %d, value %s", key, value)
	}
}
