package main

import "fmt"

func main() {

	originalNumbers := []string{"2", "3", "5", "11", "13", "1", "12", "43", "65", "666", "192"}
	//                 0  1    2     3     4     5    6     7     8     9   10

	for index_number, index_value := range originalNumbers {

		fmt.Println(index_number, "=", index_value+"**.000")
	}
}
