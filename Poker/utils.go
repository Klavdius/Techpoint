package main

import (
	"strconv"
	"strings"
)

var outMask = map[string]int{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func outMaskCard(hand string) int {
	var outMask = map[string]int{
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	cards := strings.Split(hand, " ")
	firstCard, err := strconv.Atoi(cards[0])
	if err != nil {
		firstCard = outMask[cards[0][:1]]
	}
	secondCard, err := strconv.Atoi(cards[1])
	if err != nil {
		secondCard = outMask[cards[1][:1]]
	}
	var topCard int
	if firstCard > secondCard {
		topCard = firstCard
	} else {
		topCard = secondCard
	}
	return topCard
}

func MaskCard(number int, suit string) string {
	var highTir = map[int]string{
		10: "T",
		11: "J",
		12: "Q",
		13: "K",
		14: "A",
	}
	var newCard string
	if number < 10 {
		newCard = strconv.Itoa(number) + suit
	} else {
		newCard = highTir[number] + suit
	}
	return newCard
}
