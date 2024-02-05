package main

import (
	"fmt"
	"strconv"
)

var finalResult = []int{}
var (
	sum       = 0
	sumNeg    = 0
	firstElem int
	elem      int
	nextElem  int
)

func main() {
	var amountRecord int

	fmt.Scan(&amountRecord)
	for i := 0; i < amountRecord; i++ {
		var lengthInputString int
		fmt.Scan(&lengthInputString)
		for inner := 0; inner < lengthInputString; inner++ {
			fmt.Scan(&nextElem)
			if inner != 0 {
				if elem-nextElem == 1 || elem-nextElem == -1 {
					if elem-nextElem == 1 {
						if sum == 0 {
							sumNeg--
						} else {
							AddNumberInResult(firstElem, sum)
							NewRow()
						}
						elem = nextElem
					} else {
						if sumNeg == 0 {
							sum++
						} else {
							AddNumberInResult(firstElem, sumNeg)
							NewRow()
						}
						elem = nextElem
					}
				} else {
					ChekingNeedAddNumberInRow()
					NewRow()
				}

			} else {
				NewRow()
			}

			if inner == lengthInputString-1 {
				ChekingNeedAddNumberInRow()
				NewRow()
			}
		}
		PrintResult()
	}
}

func AddNumberInResult(firstNum int, sum int) {
	finalResult = append(finalResult, firstNum)
	finalResult = append(finalResult, sum)
}

func PrintResult() {
	fmt.Println(len(finalResult))
	for _, v := range finalResult {
		fmt.Print(strconv.Itoa(v) + " ")
	}
	fmt.Println()
	finalResult = nil
	sum = 0
	sumNeg = 0
}

func ChekingNeedAddNumberInRow() {
	if sum == 0 && sumNeg == 0 {
		AddNumberInResult(firstElem, 0)
	} else if sum == 0 {
		AddNumberInResult(firstElem, sumNeg)
	} else {
		AddNumberInResult(firstElem, sum)
	}
}

func NewRow() {
	firstElem = nextElem
	elem = nextElem
	sum = 0
	sumNeg = 0
}
