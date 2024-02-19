package main

import (
	"strconv"
)

type Player struct {
	number           int
	firstCard        string
	secondCard       string
	havePocketPair   bool
	cardsNeededToWin []string
}

func (p *Player) SortCards() {
	firstNumber, err := strconv.Atoi(p.firstCard[:1])
	if err != nil {
		firstNumber = outMask[p.firstCard[:1]]
	}
	secondNumber, err := strconv.Atoi(p.secondCard[:1])
	if err != nil {
		secondNumber = outMask[p.secondCard[:1]]
	}
	if firstNumber < secondNumber {
		cup := p.firstCard
		p.firstCard = p.secondCard
		p.secondCard = cup
	}
}

func (p *Player) LookForPair() {
	if p.firstCard[:1] == p.secondCard[:1] {
		p.havePocketPair = true
	}
}

func (p *Player) FoundWinCard() {
	if !p.havePocketPair {
		for _, v := range suits {
			if v != p.firstCard[1:] {
				p.cardsNeededToWin = append(p.cardsNeededToWin, p.firstCard[:1]+v)
			}
			if v != p.secondCard[1:] {
				p.cardsNeededToWin = append(p.cardsNeededToWin, p.secondCard[:1]+v)
			}
		}
	} else {
		for _, v := range suits {
			if v == p.secondCard[1:] || v == p.firstCard[1:] {

			} else {
				p.cardsNeededToWin = append(p.cardsNeededToWin, p.firstCard[:1]+v)
			}
		}
	}
}

func (p *Player) Comparison(rival Player) {
	if IntNumberCard(rival.firstCard[:1]) > IntNumberCard(p.firstCard[:1]) {
		var cut = []string{}
		p.cardsNeededToWin = cut
	} else {
		if IntNumberCard(rival.firstCard[:1]) > IntNumberCard(p.secondCard[:1]) {
			var cut = []string{}
			for _, v := range p.cardsNeededToWin {
				if v[:1] != p.secondCard[:1] {
					cut = append(cut, v)
				}
			}
			p.cardsNeededToWin = cut
		}
	}
}
