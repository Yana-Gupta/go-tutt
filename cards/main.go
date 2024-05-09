package main

import "fmt"

func main() {
	cards := newCard()
	fmt.Println(cards)
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, ("Six of Spades"))

	// Iterate over a set or slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	cards := deck{"Ace of Diamonds", newCard()}

	cards.print()

	cards := newDeck()

	for i, card := range cards {
		fmt.Println(i, card)
	}

	deck1, cards2 := deal(cards, 5)

	for i, card := range deck1 {
		fmt.Println(i, card)
	}

	for i, card := range cards2 {
		fmt.Println(i, card)
	}

	stringDeck := cards.convertDeckToString()
	fmt.Println((stringDeck))

	cards.saveToFile("my_deck")
	my_deck := newDeckFromFile("my_deck")

	my_deck.shuffle()
	my_deck.print()
}
