package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		amountRecord, players         int
		dropFirstCard, dropSecondCard string
	)

	fmt.Fscan(in, &amountRecord)
	for i := 0; i < amountRecord; i++ {
		fmt.Fscan(in, &players)
		cardsInGame := map[int]string{}
		for i := 0; i < players; i++ {
			_, err := fmt.Fscan(in, &dropFirstCard, &dropSecondCard)
			if err != nil {
				fmt.Fscan(in, &dropFirstCard, &dropSecondCard)
			}
			cardsInGame[i] = dropFirstCard + " " + dropSecondCard
		}

		var winningCards = []string{}
		therePocketPair := ChekingPocketPair(cardsInGame)
		if therePocketPair {
			winningCards = SearchForSetWinners(cardsInGame)
		} else {
			winningCards = SearchForPairWinners(cardsInGame)
		}

		if len(winningCards) == 0 {
			winningCards = SearchHighCard(cardsInGame)
		}

		for _, v := range winningCards {
			fmt.Fprint(out, v+"\r\n")
		}
	}
}

func ChekingPocketPair(cardInGame map[int]string) bool {
	pair := false
	takeCards := strings.Split(cardInGame[0], " ")
	if takeCards[0][:1] == takeCards[1][:1] {
		pair = true
	}

	return pair
}

func SearchForSetWinners(cardsInGame map[int]string) []string {
	var suits = []string{"S", "C", "D", "H"}

	var winningCards = []string{}
	for i := 0; i < 4; i++ {
		suit := suits[i]
		for inner := 2; inner < 15; inner++ {
			newCard := MaskCard(inner, suit)
			double := FoundDuplicateCard(newCard, cardsInGame)
			if !double {
				var winners = []int{}
				winners = FoundWinnersWithSet(newCard, cardsInGame)
				for _, v := range winners {
					if v == 0 {
						winningCards = append(winningCards, newCard)
					}
				}
			}
		}
	}

	return winningCards
}

func SearchHighCard(cardsInGame map[int]string) []string {
	var winningCards = []string{}
	highCardInPocket := outMaskCard(cardsInGame[0])
	if 14-highCardInPocket == 0 {
		fmt.Println("all in!!!!!!!!!!!!!!!!!!!")
	}
	for i := 1; i < len(cardsInGame); i++ {

	}
	return winningCards
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

func SearchForPairWinners(cardsInGame map[int]string) []string {
	var suits = []string{"S", "C", "D", "H"}

	var winningCards = []string{}
	for i := 0; i < 4; i++ {
		suit := suits[i]
		for inner := 2; inner < 15; inner++ {
			newCard := MaskCard(inner, suit)
			double := FoundDuplicateCard(newCard, cardsInGame)
			if !double {
				var winners = []int{}
				winners = FoundWinnersWithPair(newCard, cardsInGame)
				for _, v := range winners {
					if v == 0 {
						winningCards = append(winningCards, newCard)
					}
				}
			}
		}
	}

	return winningCards
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

func FoundDuplicateCard(card string, cardInGame map[int]string) bool {
	result := false

	for _, v := range cardInGame {
		takeCards := strings.Split(v, " ")
		if card == takeCards[0] {
			result = true
		} else {
			if card == takeCards[1] {
				result = true
			}
		}
	}

	return result
}

func FoundWinnersWithSet(card string, cardInGame map[int]string) []int {
	var winners = []int{}
	for i, v := range cardInGame {
		takeCards := strings.Split(v, " ")
		if card[:1] == takeCards[0][:1] && card[:1] == takeCards[1][:1] {
			winners = append(winners, i)
		}
	}

	return winners
}

func FoundWinnersWithPair(card string, cardInGame map[int]string) []int {
	var winners = []int{}
	for i, v := range cardInGame {
		takeCards := strings.Split(v, " ")
		if card[:1] == takeCards[0][:1] || card[:1] == takeCards[1][:1] {
			winners = append(winners, i)
		}
	}

	return winners
}
