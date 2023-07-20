package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace,Suit: Heart})

	//Output:
	//Ace of Hearts
}

func TestNew(t *testing.T){
	cards := New()
	if len(cards) != 13*4{
		t.Error("Uneven numbers of card for the new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(Sort(Less))
	first := Card{Rank: Ace, Suit: Spade}
	if cards[0] != first{
		t.Error("Expected Ace of Spades as the first card. Received: ",cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _,c := range cards{
		if c.Suit == Joker{
			count++
		}
	}
	if count != 3{
		t.Error("Expected three jokers in the Deck. Found: ",count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(c Card)bool{
		return c.Rank == Two || c.Rank == Three
	}

	cards := New(Filter(filter))

	for _,v := range cards{
		if v.Rank == Two || v.Rank == Three{
			t.Error("Expected two and three ranks to be filtered out")
		}
	}
}

func TestDeck(t *testing.T){
	cards := New(Deck(2))
	if len(cards) != 13*4*2{
		t.Errorf("Expected %d cards but received %d cards",13*4*2,len(cards))
	}
}