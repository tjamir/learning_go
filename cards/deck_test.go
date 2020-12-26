package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected lenght was 52, but got %v", len(d))

	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card was Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card was King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected lenght was 52, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
