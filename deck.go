package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

// Create a new type deck which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Receiver type argument (equivalent to receiving "this", automatically)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Arguments as function arguments
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Receiver type argument (equivalent to receiving "this", automatically)
func (d deck) toString () string {
	return strings.Join([]string(d), ",")
}

//Receiver type argument (equivalent to receiving "this", automatically)
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}