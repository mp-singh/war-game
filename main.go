package main

import (
	"github.com/mp-singh/war/models"
	"log"
	"math/rand"
)

func main() {
	d := models.NewDeck()
	d.Shuffle()

	p1 := models.Player{Deck: models.Deck{Cards: d.Cards[:26]}}
	p2 := models.Player{Deck: models.Deck{Cards: d.Cards[26:]}}

	count := 0
	for {
		if p1.Empty() {
			if len(p1.Discard) == 0 {
				log.Printf("Player 2 Won in %d iterations!", count)
				//log.Println(len(p2.Deck.Cards) + len(p2.Discard))
				break
			} else {
				p1.Deck.Cards = p1.Discard
				p1.Deck.Shuffle()
				p1.Discard = nil
			}
		}

		if p2.Empty() {
			if len(p2.Discard) == 0 {
				log.Printf("Player 1 Won in %d iterations!", count)
				//log.Println(len(p1.Deck.Cards) + len(p1.Discard))
				break
			} else {
				p2.Deck.Cards = p2.Discard
				p2.Deck.Shuffle()
				p2.Discard = nil
			}
		}

		p1Card := p1.Draw()
		p2Card := p2.Draw()

		if p1Card.Val == p2Card.Val {
			randNum := rand.Intn(2)
			if randNum == 0 {
				p1.Discard = append(p1.Discard, []models.Card{p1Card, p2Card}...)
				log.Println("P1 won round")
			} else {
				p2.Discard = append(p2.Discard, []models.Card{p1Card, p2Card}...)
				log.Println("P2 won round")
			}

		} else if p1Card.Val > p2Card.Val {
			p1.Discard = append(p1.Discard, []models.Card{p1Card, p2Card}...)
			log.Println("P1 won round")
		} else {
			p2.Discard = append(p2.Discard, []models.Card{p1Card, p2Card}...)
			log.Println("P2 won round")
		}
		count++
	}
}
