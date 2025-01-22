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
	//srt := "hello, world"
	//fmt.Println(strings.HasPrefix(srt, "hello"))
	//fmt.Println(strings.HasSuffix(srt, "world"))
	//srt := "hello, world"
	//fmt.Println(strings.Index(srt, "world"))
	//srt := []string{"azam", "atarod"}
	//srt2 := "*"
	//fmt.Println(strings.Join(srt, srt2))
	//srt := "a, b, c, d, e, f, g, h, i, i"
	//subsrt := ","
	//fmt.Println(strings.Split(srt, subsrt))
	srt := "hello, world"
	old := "world"
	new := "Iran"
	fmt.Println(strings.Replace(srt, old, new, -1))
}
