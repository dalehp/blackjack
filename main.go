package main

import (
	"fmt"

	"github.com/user/dalehp/deck"
)

type Hand []deck.Card

func draw(d []deck.Card) (deck.Card, []deck.Card) {
	return d[0], d[1:]
}

func printDealer(h Hand) string {
	return fmt.Sprintf("------ %s", h[1:])
}

func score(h Hand) (total int) {
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
	for input != "s" {
		input = ""
		if score(player) > 21 {
			fmt.Println("You have bust. Dealer wins :(")
			return
		}
		fmt.Printf("Dealer has %s\n", printDealer(dealer))
		fmt.Printf("You have %s (%d), (s)tand or (h)it?", player, score(player))
		fmt.Scanln(&input)
		if input == "h" {
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	for ds := score(dealer); ds < 17; {
		card, cards = draw(cards)
		dealer = append(dealer, card)
		ds = score(dealer)
	}

	dp := score(player)
	ds := score(dealer)
	switch {
	case ds > 21:
		fmt.Println("Dealer busts, you win!")
	case dp > ds:
		fmt.Printf("You have %d, dealer has %d. You win!", dp, ds)
	case ds > dp:
		fmt.Printf("You have %d, dealer has %d. You lose :(", dp, ds)
	case ds == dp:
		fmt.Printf("It's a tie :/")
	}

	fmt.Println(player, dealer)
}
