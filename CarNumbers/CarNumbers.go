package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var amountRepit int
	fmt.Scanln(&amountRepit)
	scaner := bufio.NewScanner(os.Stdin)
	for i := 0; i < amountRepit; i++ {
		scaner.Scan()
		inputString := scaner.Text()
		ChekingLengInputLine(inputString)
	}
}

func ChekingLengInputLine(inputString string) {
	lengInputString := len(inputString)
	if lengInputString >= 4 {
		ChekingInputLine(inputString)
	} else {
		fmt.Println("-")
	}
}

func ChekingInputLine(inputString string) {
	var goodCarNumber = []string{}
	for len(inputString) > 0 {
		if len(inputString) < 4 {
			goodCarNumber = append(goodCarNumber, "-")
			break
		} else {
			lineForTest := inputString[:4]
			applicant := DetailFourLine(lineForTest)
			if applicant == "-" {
				if len(inputString) < 5 {
					goodCarNumber = append(goodCarNumber, "-")
					break
				} else {
					lineForTest = inputString[:5]
					applicant = DetailFiveLine(lineForTest)
					if applicant == "-" {
						goodCarNumber = append(goodCarNumber, "-")
						break
					} else {
						goodCarNumber = append(goodCarNumber, applicant)
						inputString = inputString[5:]
					}
				}
			} else {
				goodCarNumber = append(goodCarNumber, applicant)
				inputString = inputString[4:]
			}
		}
	}

	if len(inputString) != 0 {
		fmt.Println("-")
	} else {
		CheckSliceGoodCarNumber(goodCarNumber)
	}
}
func CheckSliceGoodCarNumber(inputDataInSlice []string) {
	errorInSlice := 0
	for _, v := range inputDataInSlice {
		if v == "-" {
			errorInSlice = 1
		}
	}
	if errorInSlice > 0 {
		fmt.Println("-")
	} else {
		for _, v := range inputDataInSlice {
			fmt.Print(v + " ")
		}
		fmt.Print("\n")
	}

}

func DetailFourLine(line string) string {
	var result string
	firstIndexNumber := strings.IndexAny(line, "0123456789")
	lastIndexNumber := strings.LastIndexAny(line, "0123456789")
	if firstIndexNumber == lastIndexNumber && firstIndexNumber == 1 {
		result = line
	} else {
		result = "-"
	}

	return result
}

func DetailFiveLine(line string) string {
	var result string
	firstIndexNumber := strings.IndexAny(line, "0123456789")
	lastIndexNumber := strings.LastIndexAny(line, "0123456789")
	if firstIndexNumber == 1 && lastIndexNumber == 2 {
		result = line
	} else {
		result = "-"
	}

	return result
}
