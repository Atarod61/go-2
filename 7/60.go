package main

import "log"

func main() {
	isError := true
	if isError {
		log.Println("we have error here")
		log.Printf("your Error statues: %d", 404)
	}

}
