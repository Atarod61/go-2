package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"os/user"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter somthing: ")
	scanner.Scan()
	userlnput := scanner.Text()
	//Try to parse the user input as a number
	if num, err := strconv.Atoi(userlnput) ; err == nil{
		fmt.Printf("%d is an integer\n", num)
	}else if -, err := strconv.ParseFloat(userIuserInput, 64); err == nil {
		fmt.Printf("%s float\n", userIuserInput)
	}else{
		fmt.Printf("%s string\n", userIuserInput)
	}
}
