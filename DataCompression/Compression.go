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
		amountRecords   int
		lengInputString int
		inputString     string
	)

	finalLine := []int{}

	fmt.Scanln(&amountRecords)
	for i := 0; i < amountRecords; i++ {
		fmt.Scanln(&lengInputString)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputString = scanner.Text()
		inputRawData := strings.Split(inputString, " ")
		inputData := ConvertSliceToInt(inputRawData)

		for inner := 0; inner < lengInputString; inner++ {

			lengTestLine := FindLengRowInInput(inputData)

			testOnZero, positionZero := TestOnDoubleZeroInLine(inputData[:lengTestLine])

			for testOnZero > 1 {
				lengTestLine = CuteDoubleZeroInLine(inputData[:positionZero+1])
				testOnZero, positionZero = TestOnDoubleZeroInLine(inputData[:lengTestLine])
			}

			test2 := testOnOneRowSing(inputData[:lengTestLine])

			for !test2 {
				lengTestLine--
				test2 = testOnOneRowSing(inputData[:lengTestLine])
			}
			line := inputData[:lengTestLine]
			sing := FindRowSing(line)

			number := 0
			switch sing {
			case "":
				number = 0
			case "-":
				number = (len(line) - 1) * -1
			case "+":
				number = (len(line) - 1)
			}
			finalLine = append(finalLine, line[0])
			finalLine = append(finalLine, number)
			inputData = inputData[len(line):]

			if len(inputData) == 0 {
				break
			}
		}
		fmt.Println(len(finalLine))

		for _, v := range finalLine {
			fmt.Print(strconv.Itoa(v) + " ")
		}
		fmt.Println()
		finalLine = nil
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

func FindLengRowInInput(line []int) int {
	firstElem := line[0]
	nextElem := firstElem
	sum := 1
	for i, v := range line {
		if i != 0 {
			if nextElem-v == 1 || nextElem-v == -1 {
				sum++
				nextElem = v
			} else {
				break
			}
		}
	}

	return sum
}

func TestOnDoubleZeroInLine(line []int) (int, int) {
	result := 0
	positionZero := 0
	for i, v := range line {
		if v == 0 {
			result++
			positionZero = i
		}
	}

	return result, positionZero
}

func CuteDoubleZeroInLine(line []int) int {
	line = line[:len(line)-1]
	return len(line)
}

func testOnOneRowSing(line []int) bool {
	result := true
	if len(line) != 1 {
		sing := line[0] - line[1]
		if sing > 0 {
			for i, v := range line {
				if i != 0 {
					if v >= line[0] {
						result = false
					}
				}
			}
		} else {
			for i, v := range line {
				if i != 0 {
					if v <= line[0] {
						result = false
					}
				}
			}
		}
	}

	return result
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
