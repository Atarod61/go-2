package main

import (
	"fmt"
	"strings"
)

func main() {

	//srt := "hello, world"

	//subsrt := "world"
	//fmt.Println(strings.Contains(srt, subsrt))
	//srt := "banana"
	//subsrt := "a"
	//fmt.Println(strings.Count(srt, subsrt))
	srt := "hello, world"
	fmt.Println(strings.HasPrefix(srt, "hello"))
	fmt.Println(strings.HasSuffix(srt, "world"))

}
