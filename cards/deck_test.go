package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expected := 52
	if len(d) != expected {
		t.Errorf("Expected 52, but got %v", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("Expected ace of spades, but got %v", d[0])
	}

	if d[len(d)-1] != "king of clubs" {
		t.Errorf("Expected ace of spades, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
