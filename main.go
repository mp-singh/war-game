package main

import (
	"github.com/mp-singh/war-game/models"
	"log"
	"math/rand"
)

func main() {
	d := models.NewDeck()
	d.Shuffle()

	p1 := models.Player{Name: "Player One", Deck: models.Deck{Cards: d.Cards[:26]}}
	p2 := models.Player{Name: "Player Two", Deck: models.Deck{Cards: d.Cards[26:]}}

	count := 0
	for {
		if p1Won := checkWinCondition(p1, count); p1Won {
			break
		}

		if p2Won := checkWinCondition(p2, count); p2Won {
			break
		}

		p1Card := p1.Draw()
		p2Card := p2.Draw()

		if p1Card.Val == p2Card.Val {
			randNum := rand.Intn(2)
			if randNum == 0 {
				p1.Discard = append(p1.Discard, []models.Card{p1Card, p2Card}...)
				log.Printf("%s won round!", p1.Name)
			} else {
				p2.Discard = append(p2.Discard, []models.Card{p1Card, p2Card}...)
				log.Printf("%s won round!", p2.Name)
			}

		} else if p1Card.Val > p2Card.Val {
			p1.Discard = append(p1.Discard, []models.Card{p1Card, p2Card}...)
			log.Printf("%s won round!", p1.Name)
		} else {
			p2.Discard = append(p2.Discard, []models.Card{p1Card, p2Card}...)
			log.Printf("%s won round!", p2.Name)
		}
		count++
	}
}

func checkWinCondition(p models.Player, count int) bool {
	if p.Empty() {
		if len(p.Discard) == 0 {
			log.Printf("%s WON in %d iterations!", p.Name, count)
			return true
		} else {
			p.Deck.Cards = p.Discard
			p.Deck.Shuffle()
			p.Discard = nil
			return false
		}
	}
	return false
}
