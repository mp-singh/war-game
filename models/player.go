package models

type Player struct {
	Name string
	Deck    Deck
	Discard []Card
}

func (p *Player) Empty() bool {
	return len(p.Deck.Cards) == 0
}


func (p *Player) Draw() Card {
	card, cards := p.Deck.Cards[0], p.Deck.Cards[1:]
	p.Deck.Cards = cards
	return card
}