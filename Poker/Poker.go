package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var (
			playersList = []Player{}
			cardsInGame = map[string]string{}
		)
		cardsInGame = BuildNewDeck()
		fmt.Fscan(in, &players)
		var c Catcher
		var p Player
		for i := 0; i < players; i++ {
			_, err := fmt.Fscan(in, &dropFirstCard, &dropSecondCard)
			if err != nil {
				fmt.Fscan(in, &dropFirstCard, &dropSecondCard)
			}
			c.number = i
			c.firstCard = dropFirstCard
			c.secondCard = dropSecondCard
			p = MakePlayer(c)
			playersList = append(playersList, p)
			delete(cardsInGame, dropFirstCard)
			delete(cardsInGame, dropSecondCard)
		}

		if !playersList[0].havePocketPair {
			for _, v := range playersList {
				if v.havePocketPair == true {
					playersList[0].Comparison(v)
				}
			}
			if len(playersList[0].cardsNeededToWin) == 0 {
				fmt.Println(0)
			} else {
				playersList[0].cardsNeededToWin = FindWinCardInDeck(cardsInGame, playersList[0].cardsNeededToWin)
				if len(playersList[0].cardsNeededToWin) == 0 {
					TopCardInHand(playersList, cardsInGame)
				} else {
					TopCardInHand(playersList, cardsInGame)
					//PrintWin(playersList[0].cardsNeededToWin)
				}

			}
		} else {
			WinToPocketPair(playersList, cardsInGame)
			//PrintWin(playersList[0].cardsNeededToWin)
		}

	}

}

func MakePlayer(c Catcher) Player {
	var p Player
	p.number = c.number
	p.firstCard = c.firstCard
	p.secondCard = c.secondCard
	p.SortCards()
	p.LookForPair()
	p.FoundWinCard()
	return p
}

func PrintWin(line []string) {
	sort.Strings(line)
	fmt.Println(len(line))
	for _, v := range line {
		fmt.Println(v)
	}
}

func TopCardInHand(list []Player, deck map[string]string) {
	for _, v := range list {
		for _, card := range v.cardsNeededToWin {
			_, ok := deck[card]
			if ok {
				delete(deck, card)
			}
		}
	}
	fmt.Println(len(deck))
	var key = []string{}
	for _, v := range deck {
		key = append(key, v)
	}
	sort.Strings(key)
	for _, v := range key {
		fmt.Println(v)
	}
}

func WinToPocketPair(list []Player, deck map[string]string) {
	for _, v := range list {
		for _, card := range v.cardsNeededToWin {
			_, ok := deck[card]
			if ok {
				if v.havePocketPair {
					delete(deck, card)
				} else {
					if IntNumberCard(card[:1]) > IntNumberCard(list[0].firstCard[:1]) {
						delete(deck, card)
					}
				}
			}
		}
	}

	fmt.Println(len(deck))
	var key = []string{}
	for _, v := range deck {
		key = append(key, v)
	}
	sort.Strings(key)
	for _, v := range key {
		fmt.Println(v)
	}
}
