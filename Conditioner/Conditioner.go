package main

import (
	"fmt"
	"strconv"
	"strings"
)

var band = []int{15, 30}

func main() {
	var (
		amount       int
		amountPeople int
		inputSign    string
		inputNumber  string
	)

	checkErrorBand := false
	fmt.Scanln(&amount)
	for i := 0; i < amount; i++ {
		fmt.Scan(&amountPeople)
		for inner := 0; inner < amountPeople; inner++ {
			fmt.Scan(&inputSign)
			checkChar := strings.EqualFold(inputSign, "<=")
			fmt.Scan(&inputNumber)
			if checkErrorBand {
				fmt.Println("-1")
			} else {
				if checkChar {
					checkErrorBand = CheckingNumberNoMoreMax(inputNumber)
				} else {
					checkErrorBand = CheckingNumberNoLessMin(inputNumber)
				}

				if checkErrorBand {
					fmt.Println("-1")
				} else {
					if checkChar {
						ChangingRadiusBoundariesMax(inputNumber)
					} else {
						ChangingRadiusBoundariesMin(inputNumber)
					}
					fmt.Println(strconv.Itoa(band[0]))
				}
			}
		}
		band[0] = 15
		band[1] = 30
		checkErrorBand = false
	}
}

func CheckingNumberNoMoreMax(inputText string) bool {
	inputNumber, _ := strconv.Atoi(inputText)
	var result bool
	if inputNumber < band[0] {
		result = true
	} else {
		result = false
	}
	return result
}

func CheckingNumberNoLessMin(inputText string) bool {
	inputNumber, _ := strconv.Atoi(inputText)
	var result bool
	if inputNumber > band[1] {
		result = true
	} else {
		result = false
	}
	return result
}

func ChangingRadiusBoundariesMax(input string) {
	limit, _ := strconv.Atoi(input)
	if limit < band[1] {
		band[1] = limit
	}
}

func ChangingRadiusBoundariesMin(input string) {
	limit, _ := strconv.Atoi(input)
	if limit > band[0] {
		band[0] = limit
	}
}
