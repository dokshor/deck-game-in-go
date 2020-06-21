package main

import (
	"os"
	"testing"
)

func TestNewDesk(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but we got %v", len(d))
	}
}

func TestNewDeskFirstElement(t *testing.T) {
	d := newDeck()

	if d[0] != "One of Spades" {
		t.Errorf("Expected One of Spades as the first element in the []deck but we got %v", d[0])
	}
}

func TestNewDeskLastElement(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs as the last element in the []deck but we got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewTestFromFile(t *testing.T) {
	os.Remove("cards.db.test")

	if _, err := os.Stat("file-exists.go.test"); err == nil {
		t.Errorf("Cards database file exists and has not been removed as file-exists.go.test")
	}

	d := newDeck()
	d.saveToFile("cards.db.test")

	if _, err := os.Stat("file-exists.go.test"); err == nil {
		t.Errorf("Cards database file has not been created as file-exists.go.test")
	}

	loadDeck := newDeckFromFile("cards.db.test")

	if len(loadDeck) != 16 {
		t.Errorf("Expected deck database file cards.db.test length of 16 but we got %v", len(loadDeck))
	}

	os.Remove("cards.db.test")
}
