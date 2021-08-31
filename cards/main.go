package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards.txt")
	newCards := newDeckFromFile("my_cards.txt")
	newCards.print()

}
