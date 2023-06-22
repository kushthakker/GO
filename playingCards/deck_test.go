package main
import ("testing"
"os")


func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected deck length of 20, but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card of Ace of spades, but got %v", d[0])
	}

	if d[len(d)-1] != "four of Clubs" {
		t.Errorf("expected lasr card of Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndnewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToHardrive("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards, got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}