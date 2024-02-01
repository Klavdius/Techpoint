package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		amountRecords         int
		lengInputString       int
		previousElem          int
		rowSing               string
		inputString           string
		lineNotHaveZeroNumber bool
		flag                  bool
	)
	fmt.Scanln(&amountRecords)
	for i := 0; i < amountRecords; i++ {
		fmt.Scanln(&lengInputString)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputString = scanner.Text()
		inputRawData := strings.Split(inputString, " ")
		inputData := ConvertSliceToInt(inputRawData)

		firstElem := inputData[0]
		previousElem = inputData[0]
		lineNotHaveZeroNumber = true
		line := []int{}
		for inner := 1; inner < lengInputString; inner++ {
			if lineNotHaveZeroNumber == true {
				if inputData[inner] == 0 {
					flag = true
				}
			} else {
				if inputData[inner] == 0 {
					rowSing = FindRowSing(line)
					fmt.Print(strconv.Itoa(firstElem) + " " + rowSing + strconv.Itoa(len(line)) + " ")
					firstElem = inputData[inner]
					line = nil
					lineNotHaveZeroNumber = true
					flag = false
				}

			}

			if previousElem-inputData[inner] == 1 || previousElem-inputData[inner] == -1 {
				line = append(line, inputData[inner])
				if flag {
					lineNotHaveZeroNumber = false
				}
			} else {
				rowSing = FindRowSing(line)
				fmt.Print(strconv.Itoa(firstElem) + " " + rowSing + strconv.Itoa(len(line)) + " ")
				firstElem = inputData[inner]
				line = nil
				lineNotHaveZeroNumber = true
				flag = false
			}

			previousElem = inputData[inner]
		}
	}
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

func FindRowSing(line []int) string {
	var result string
	if len(line) == 1 {
		result = ""
	} else {
		if line[0] > line[1] {
			result = "-"
		} else {
			result = "+"
		}
	}
	return result
}
