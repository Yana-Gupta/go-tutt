package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heart"}

	cardValues := []string{"One", "two", "three", "four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, (suite + " " + value))
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck - []string - string - []byte
func (d deck) convertDeckToString() string {
	return (strings.Join([]string(d), ","))
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.convertDeckToString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println((err))
		os.Exit(1)
	}

	ele1 := strings.Split(string(bs), ",")

	dec := deck(ele1)

	return dec

}

func (d deck) print() {

	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) shuffle() {
	// Its not always required to give the reference to the element you are referring to

	// Go random Number generator use a seed

	// This seed is same for default

	// Create a new source to ensure randomness

	source := rand.NewSource(time.Now().UnixNano() + time.Hour.Microseconds())

	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
