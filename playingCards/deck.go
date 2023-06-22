package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"crypto/rand"
	"os"
	"strings"
)

type deck []string
// type combination string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _,suit := range cardSuits {
		for _,value := range cardValues {
			cards = append(cards, value+" of "+suit)
			// fmt.Println(value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println(d, "d")
	for i,card := range d {
		fmt.Println(i,card)
	}
}

func deal(d deck,handSize int) (deck, deck) {
	  fmt.Println(d)
	  return d[:handSize], d[handSize:]
}

func (c color) describe(description string) (string) {
	return string(c) + " " + description
 }

func (d deck) toString() string {
	// combination := ""
	// for _,ele := range d {
	// 	combination = combination + " " + ele
	// }
	// return combination
	return strings.Join([]string(d), ",")
	
}

func stringToByte(toBeConverted string) []byte {
	return []byte(toBeConverted)
}

func (d deck) saveToHardrive(name string) error {
	return ioutil.WriteFile(name, stringToByte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs,err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleDeck() {
	for i := range d {
		// newPositions := rand.Intn(len(d)-1)
		// d[i], d[newPositions] = d[newPositions], d[i]
		idx,_ := rand.Int(rand.Reader, big.NewInt(int64(len(d)-1)))
		fmt.Println(idx, "idx")
		newPosition := idx.Int64()
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}