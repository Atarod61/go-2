package main

import (
	"fmt"
	"strings"
)

func main() {

	srt := "hello, world"

	subsrt := "world"
	fmt.Println(strings.Contains(srt, subsrt))

}
