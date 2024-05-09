package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))
	}

	if d[0] != "Spades One" {
		t.Errorf("Expected first card Spades One, but got %s", d[0])
	}

	if d[len(d)-1] != "Heart four" {
		t.Errorf("Expected last card Heart Four, but got %s", d[len(d)-1])
	}
}

// Remember you have to just delete the temporary files
// You need to clean up the file yourself
// GO will not do it for you
// Steps -
// 1. delete all the files in the workdir with name _deckTesting
// 2. Create the deck
// 3. Save to the directory
// 4. Remove the directory
func TestSaveToFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")
	l2 := newDeckFromFile("_deckTesting")
	if len(l2) != 12 {
		t.Errorf("Expected 12 cards in dck, got %v cards", len(l2))
	}
	os.Remove("_deckTesting")
}
