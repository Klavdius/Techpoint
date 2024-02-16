package main

import (
	"bufio"
	"fmt"
	"os"
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
	var playersList = []Player{}
	var cardInHands = []string{}
	var winCards = []string{}

	for i := 0; i < amountRecord; i++ {
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
			cardInHands = append(cardInHands, dropFirstCard, dropSecondCard)
		}

		if !ChekingOnPocketPair(playersList[0]) {

		} else {
			NeedWinCards(&winCards, playersList[0])
			fmt.Println(winCards)
		}
	}
	//fmt.Println(playersList)
}

func MakePlayer(c Catcher) Player {
	var p Player
	p.number = c.number
	p.firstCard = c.firstCard
	p.secondCard = c.secondCard
	p.SortCards()
	p.LookForPair()
	return p
}

func NeedWinCards(line *[]string, p Player) {
	numberCards := p.firstCard[:1]
	var cup = []string{}
	for _, v := range fats {
		if v != p.firstCard[1:] {
			if v != p.secondCard[1:] {
				cup = append(cup, numberCards+v)
			}
		}
	}
	line = append(line, cup[0], cup[1])
}
