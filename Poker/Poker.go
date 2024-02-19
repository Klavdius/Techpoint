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
			CardsInGame = map[string]string{}
		)
		CardsInGame = BuildNewDeck()
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
			delete(CardsInGame, dropFirstCard)
			delete(CardsInGame, dropSecondCard)
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
				PrintWin(playersList[0].cardsNeededToWin)
			}
		} else {
			PrintWin(playersList[0].cardsNeededToWin)
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
	for _, v := range line {
		fmt.Println(v)
	}
}
