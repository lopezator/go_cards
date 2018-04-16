package main

func main() {
	//1.- Deck to file
	//deck := newDeck()
	//deck.saveToFile("my_cards")

	//2.- Deck from file
	deck := newDeckFromFile("my_cards")
	deck.shuffle()
	deck.print()
}