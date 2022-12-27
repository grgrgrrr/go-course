package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck lenght of 20 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first value to be Ace od Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected the first value to be Ace od Spades, but got %v", d[0])
	}
}

func TestDeckSave(t *testing.T) {

}
