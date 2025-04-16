package main

import "log"

func main() {
	isError := true
	if isError {
		//log.Println("we have error here")
		//log.Printf("your Error statues: %d", 404)
		log.SetPrefix("ErrorMessage:")
		log.SetFlags(log.Ldate | log.Ltime)
		log.Printf("your Error statues: %d", 404)
	} else {
		log.Println("your program was started successfully")
	}

}
