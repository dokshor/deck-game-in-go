package main

func main() {
	cards := newDeckFromFile("cards.db")
	cards.shuffle()
	cards.print()
}
