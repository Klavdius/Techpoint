package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var amountDoc int
	fmt.Scan(&amountDoc)
	for i := 0; i < amountDoc; i++ {

		var in *bufio.Reader
		var out *bufio.Writer
		in = bufio.NewReader(os.Stdin)
		out = bufio.NewWriter(os.Stdout)

		var (
			needPage  int
			inputData string
		)
		fmt.Fscan(in, &needPage, &inputData)

		needPrintPage := []int{}
		inputPage := []string{}
		for i := 1; i <= needPage; i++ {
			needPrintPage = append(needPrintPage, i)
		}

		if len(inputData) != 1 {
			inputPage = strings.Split(inputData, ",")
		} else {
			inputPage = append(inputPage, inputData)
		}
		page := UnpackingRow(inputPage)
		needPrintPage = SelectionPage(needPrintPage, page)
		sort.Ints(needPrintPage)
		outputPage := GiveFormatRowPage(needPrintPage)
		finalString := strings.Join(outputPage, ",")
		fmt.Fprintln(out, finalString)
		out.Flush()
	}
}

func SelectionPage(needPage []int, inputPage []int) []int {
	result := []int{}
	for _, v := range needPage {
		double := false
		for _, value := range inputPage {
			if v == value {
				double = true
			}
		}
		if !double {
			result = append(result, v)
		}
	}
	return result
}

func GiveFormatRowPage(input []int) []string {
	result := []string{}
	sum := 0
	var first int
	for i, v := range input {
		if i == 0 {
			first = v
		}
		if i != len(input)-1 {
			if v-input[i+1] == -1 {
				sum++
			} else {
				if sum != 0 {
					startRow := strconv.Itoa(first)
					endRow := strconv.Itoa(first + sum)
					row := startRow + "-" + endRow
					result = append(result, row)
					first = input[i+1]
					sum = 0
				} else {
					number := strconv.Itoa(v)
					result = append(result, number)
					first = v
				}
			}
		} else {
			if sum != 0 {
				startRow := strconv.Itoa(first)
				endRow := strconv.Itoa(first + sum)
				row := startRow + "-" + endRow
				result = append(result, row)
			} else {
				number := strconv.Itoa(v)
				result = append(result, number)
			}
		}
	}

	return result
}

func UnpackingRow(input []string) []int {
	result := []int{}
	double := false
	for _, v := range input {
		solution, _ := regexp.MatchString("-", v)
		if !solution {
			number, _ := strconv.Atoi(v)
			double = false
			for _, val := range result {
				if val == number {
					double = true
				}
			}
			if !double {
				result = append(result, number)
			}
		} else {
			microRow := strings.Split(v, "-")
			startRow, _ := strconv.Atoi(microRow[0])
			endRow, _ := strconv.Atoi(microRow[1])
			for i := startRow; i <= endRow; i++ {
				double = false
				for _, val := range result {
					if val == i {
						double = true
					}
				}
				if !double {
					result = append(result, i)
				}
			}
		}
	}
	return result
}
