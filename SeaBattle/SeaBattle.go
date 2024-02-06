package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var amountLine int
	var allSlice = [][]string{}

	fmt.Scanln(&amountLine)
	for i := 0; i < amountLine; i++ {
		line := myScan()
		SliceLine := strings.Split(line, " ")
		allSlice = append(allSlice, SliceLine)
	}
	CheckSummSlice(allSlice, amountLine)
}

func myScan() string {
	newScaner := bufio.NewScanner(os.Stdin)
	newScaner.Scan()
	return newScaner.Text()
}

func CheckSummSlice(allSlice [][]string, num int) {
	for _, v := range allSlice {
		answer := CountElementInSlice(v)
		for i := 0; i < num; i++ {
			if i < (num - 1) {
				fmt.Println(answer)
			} else {
				fmt.Print(answer)
			}
		}
	}
}

func CountElementInSlice(slice []string) string {
	var (
		ship1  int
		ship2  int
		ship3  int
		ship4  int
		result string
	)
	for _, v := range slice {
		switch v {
		case "1":
			ship1 = ship1 + 1
		case "2":
			ship2 = ship2 + 1
		case "3":
			ship3 = ship3 + 1
		case "4":
			ship4 = ship4 + 1
		}
	}
	if ship4 == 1 && ship3 == 2 && ship2 == 3 && ship1 == 4 {
		result = "YES"
	} else {
		result = "NO"
	}
	return result
}
