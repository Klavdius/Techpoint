package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		amountRecords   int
		lengInputString int
		inputString     string
	)
	fmt.Scan(&amountRecords)
	for i := 0; i < amountRecords; i++ {
		fmt.Scan(&lengInputString)
		fmt.Scanln(&inputString)
		inputRawData := strings.Split(inputString, " ")
		inputData := ConvertSliceToInt(inputRawData)

		for inner := 0; inner < lengInputString; {
			step := FindRow(inputData)
		}
	}
}

func FindRow(input []int) int {
	baseElement := input[0]
	var (
		summa int
		multi int
	)

	for i := 1; i < len(input); i++ {
		nextNumber := input[i]
		switch baseElement - nextNumber {
		case 1:
			summa++
			multi = 1
		case -1:
			summa++
			multi = -1
		default:
			summa = 0
			multi = 0
		}
		baseElement = nextNumber
	}
	return 1
}

func ConvertSliceToInt(input []string) []int {
	var intSlice = []int{}
	lengSlice := len(input)
	for i := 0; i < lengSlice; i++ {
		elem, _ := strconv.Atoi(input[i])
		intSlice = append(intSlice, elem)
	}
	return intSlice
}
