package main

import (
	"fmt"
	"strings"

	"github.com/user/dalehp/deck"
)

type Hand []deck.Card

func draw(d []deck.Card) (deck.Card, []deck.Card) {
	return d[0], d[1:]
}

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i, card := range h {
		strs[i] = card.String()
	}
	return strings.Join(strs, ", ")
}

func printDealer(h Hand) string {
	return h[0].String() + ", ------"
}

/* Score returns the score of a hand */
func (h Hand) Score() (total int) {
	ace := false
	for _, c := range h {
		if c.Rank >= deck.Ten {
			total += 10
		} else {
			total += int(c.Rank)
		}
		if c.Rank == deck.Ace {
			ace = true
		}
	}
	if ace && total <= 11 {
		total += 10
	}
	return
}

func main() {
	cards := deck.New()
	cards = deck.Shuffle(cards)
	var card deck.Card
	var player, dealer Hand

	// Deal cards
	for i := 0; i < 2; i++ {
		for _, p := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*p = append(*p, card)
		}
	}

	var input string
	for input != "s" && player.Score() <= 21 {
		input = ""
		fmt.Printf("Dealer has %s\n", printDealer(dealer))
		fmt.Printf("You have %s (%d), (s)tand or (h)it?", player, player.Score())
		fmt.Scanln(&input)
		if input == "h" {
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	for ds := dealer.Score(); ds < 17; {
		card, cards = draw(cards)
		dealer = append(dealer, card)
		ds = dealer.Score()
	}

	dp := player.Score()
	ds := dealer.Score()
	fmt.Printf("DEALER: %s\n", dealer)
	fmt.Printf("PLAYER: %s\n", player)
	switch {
	case dp > 21:
		fmt.Println("Player busts, you lose.")
	case ds > 21:
		fmt.Println("Dealer busts, you win!")
	case dp > ds:
		fmt.Printf("You have %d, dealer has %d. You win!", dp, ds)
	case ds > dp:
		fmt.Printf("You have %d, dealer has %d. You lose :(", dp, ds)
	case ds == dp:
		fmt.Printf("It's a tie :/")
	}
}
