package main

import "fmt"

func main() {

	cardsAce := []string{}
	cards := []string{"Ace Pik!", "1nanbarbari!!", card()}

	//Ading to slices
	cards = append(cards, "Ace khesht!!")
	fmt.Println(cards)
	fmt.Println(cardsAce)
}
func card() string {
	return "Ace Del!"
}
