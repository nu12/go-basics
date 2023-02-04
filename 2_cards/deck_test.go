package main

import (
	"os"
	"testing"
)

const (
	DECKSIZE = 52
	TESTFILE = "_decktesting"
	HANDSIZE = 4
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != DECKSIZE {
		t.Errorf("Expected %v cards, got %v", DECKSIZE, len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card to be Ace of Spade, got %v", d[0])
	}

	if d[len(d)-1] != "King of Club" {
		t.Errorf("Expected first card to be King of Club, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove(TESTFILE)

	d := newDeck()
	d.saveToFile(TESTFILE)

	nd := readFromFile((TESTFILE))

	if len(nd) != DECKSIZE {
		t.Errorf("Expected %v cards, got %v", DECKSIZE, len(d))
	}

	os.Remove(TESTFILE)
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, d := d.deal(HANDSIZE)

	if len(hand) != HANDSIZE {
		t.Errorf("Expected %v cards, got %v", HANDSIZE, len(hand))
	}
	if len(d) != DECKSIZE-HANDSIZE {
		t.Errorf("Expected %v cards, got %v", DECKSIZE-HANDSIZE, len(hand))
	}
}
