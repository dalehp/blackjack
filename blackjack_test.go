package main

import (
	"testing"

	"github.com/user/dalehp/deck"
)

func TestScoreSimple(t *testing.T) {
	h := Hand{deck.Card{Rank: deck.Ten, Suit: deck.Club}, deck.Card{Rank: deck.Two, Suit: deck.Heart}}
	exp := 12
	act := score(h)
	if act != exp {
		t.Errorf("Hand score is incorrect, got: %d, want: %d", act, exp)
	}
}

func TestScoreAceHigh(t *testing.T) {
	h := Hand{deck.Card{Rank: deck.Ten, Suit: deck.Club}, deck.Card{Rank: deck.Ace, Suit: deck.Heart}}
	exp := 21
	act := score(h)
	if act != exp {
		t.Errorf("Hand score is incorrect, got: %d, want: %d", act, exp)
	}
}

func TestScoreAceLow(t *testing.T) {
	h := Hand{deck.Card{Rank: deck.Ten, Suit: deck.Club}, deck.Card{Rank: deck.Six, Suit: deck.Club}, deck.Card{Rank: deck.Ace, Suit: deck.Heart}}
	exp := 17
	act := score(h)
	if act != exp {
		t.Errorf("Hand score is incorrect, got: %d, want: %d", act, exp)
	}
}

func TestScoreFaceCards(t *testing.T) {
	h := Hand{deck.Card{Rank: deck.King, Suit: deck.Club}, deck.Card{Rank: deck.Queen, Suit: deck.Heart}}
	exp := 20
	act := score(h)
	if act != exp {
		t.Errorf("Hand score is incorrect, got: %d, want: %d", act, exp)
	}
}
