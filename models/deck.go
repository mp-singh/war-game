package models

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	d := Deck{}
	valEnums := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suits := []string{"Spades", "Hearts", "Diamonds", "Cloves"}
	for k, val := range valEnums {
		for _, suit := range suits {
			d.Cards = append(d.Cards, Card{
				Val:      k + 2,
				Suit:     suit,
				CardName: val,
			})
		}
	}
	return &d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}
