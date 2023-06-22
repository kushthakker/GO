package main

import "fmt"
type color string
func main(){
	// var card string = "Ace of spades" // card := "Ace of spades"
	// card := newCard()
	// fmt.Println(card)
	// cards := deck{"Ace of spades", "Ace of spades", newCard()}
	// cards = append(cards, "newly added")
	// fmt.Println(cards)
	// for i,card := range cards {
	// 	fmt.Println(i,card)
	// }
	// cards.print()

	cards := newDeck()

	hand, remaningDeck := deal(cards,5)
	hand.print()
	remaningDeck.print()

	c := color("Red")
	fmt.Println(c)

	newString := c.describe("is my color")
	fmt.Println(newString)

	// stringOutput := hand.toString()
	
	// bytedString := stringToByte(stringOutput)
	// fmt.Println(bytedString)

	// hand.saveToHardrive("test")
	newHand := newDeckFromFile("test")
	fmt.Println(newHand, "newHand")

	remaningDeck.shuffleDeck()
	remaningDeck.print()
}


// func newCard() string {
// 	return "Five of diamonds"
// } 