package main

import (
	"strconv"
)

var outMask = map[string]int{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}
var highTir = map[int]string{
	10: "T",
	11: "J",
	12: "Q",
	13: "K",
	14: "A",
}

var suits = []string{"D", "S", "C", "H"}

func IntNumberCard(line string) int {
	var number int
	number, err := strconv.Atoi(line)
	if err != nil {
		number = outMask[line]
	}
	return number
}

func TakeCard(number int, suit string) string {
	var newCard string
	if number < 10 {
		newCard = strconv.Itoa(number) + suit
	} else {
		newCard = highTir[number] + suit
	}
	return newCard
}

func BuildNewDeck() map[string]string {
	var (
		deck = map[string]string{}
		card string
	)
	for i := 2; i < 15; i++ {
		for inner := 0; inner < len(suits); inner++ {
			card = TakeCard(i, suits[inner])
			deck[card] = card
		}
	}

	return deck
}

func FindWinCardInDeck(deck map[string]string, list []string) []string {
	var winCard = []string{}
	for _, v := range list {
		_, ok := deck[v]
		if ok {
			winCard = append(winCard, v)
		}
	}

	return winCard
}
