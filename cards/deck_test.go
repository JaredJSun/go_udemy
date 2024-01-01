package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}	
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	handSize := 5

	// 分发牌
	hand, remainingDeck := deal(d, handSize)

	// 检验分发出去的牌组长度
	if len(hand) != handSize {
		t.Errorf("Expected hand size of %v, but got %v", handSize, len(hand))
	}

	// 检验剩余牌组长度
	if len(remainingDeck) != len(d)-handSize {
		t.Errorf("Expected remaining deck size of %v, but got %v", len(d)-handSize, len(remainingDeck))
	}

	// 验证是否正确分发了牌
	for i := 0; i < handSize; i++ {
		if hand[i] != d[i] {
			t.Errorf("Expected card %v in hand, but got %v", d[i], hand[i])
		}
	}
}