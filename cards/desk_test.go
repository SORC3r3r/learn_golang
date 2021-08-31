package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card in deck to be 'Four of Clubs', but got '%v'", d[len(d)-1])
	}
}

func TestDealCards(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 3)

	if len(hand) != 3 {
		t.Errorf("Expected hand to contain 3 cards, but got '%v'", len(hand))
	}

	if len(remainingDeck) != len(d)-len(hand) {
		t.Errorf("Expected hand to contain '%v' cards, but got '%v'", len(d)-len(hand), len(remainingDeck))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card in deck to be 'Four of Clubs', but got '%v'", d[len(d)-1])
	}

	os.Remove("_decktesting")
}
