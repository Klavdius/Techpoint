package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var countRepit int
	fmt.Scanln(&countRepit)
	var scaner = bufio.NewScanner(os.Stdin)
	for i := 0; i < countRepit; i++ {
		scaner.Scan()
		dateFromInput := strings.Split(scaner.Text(), " ")
		dateSlice := ConvDateSliceToInt(dateFromInput)
		ChekingDate(dateSlice)
	}
}

func ConvDateSliceToInt(data []string) []int {
	var (
		day   int
		mount int
		year  int
	)
	mount, _ = strconv.Atoi(data[1])
	day, _ = strconv.Atoi(data[0])
	year, _ = strconv.Atoi(data[2])
	var dataLine = []int{}
	dataLine = append(dataLine, day)
	dataLine = append(dataLine, mount)
	dataLine = append(dataLine, year)
	return dataLine
}

func ChekingDate(data []int) {
	var (
		checkDay   int
		checkMount int
		checkYear  int
	)
	if data[2] >= 1950 && data[2] <= 2300 {
		checkYear = 1
	} else {
		checkYear = 0
	}

	if data[1] >= 1 && data[1] <= 12 {
		checkMount = 1
	} else {
		checkMount = 0
	}

	switch data[1] {
	case 2:
		checkDay = ChekingLeapYear(data[0], data[2])
	case 1, 3, 5, 7, 8, 10, 12:
		checkDay = CheckDayLimit31(data[0])
	case 4, 6, 9, 11:
		checkDay = CheckDayLimit30(data[0])
	}

	if checkDay == 1 && checkMount == 1 && checkYear == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func ChekingLeapYear(day int, year int) int {
	var (
		checkDay int
		par1     int
		par2     int
	)
	par1 = Parameter1(year)
	par2 = Parameter2(year)

	if par1 == 1 || par2 == 1 {
		checkDay = CheckDayLeapYear(day)
	} else {
		checkDay = CheckDayNoLeapYear(day)
	}

	return checkDay
}

func Parameter2(year int) int {
	if year%400 == 0 {
		return 1
	} else {
		return 0
	}
}

func Parameter1(year int) int {
	if year%4 == 0 && year%100 != 0 {
		return 1
	} else {
		return 0
	}
}
func CheckDayLeapYear(dataDay int) int {
	var result int
	if dataDay >= 1 && dataDay <= 29 {
		result = 1
	} else {
		result = 0
	}
	return result
}

func CheckDayNoLeapYear(dataDay int) int {
	var result int
	if dataDay >= 1 && dataDay <= 28 {
		result = 1
	} else {
		result = 0
	}
	return result
}
func CheckDayLimit31(dataDay int) int {
	var result int
	if dataDay >= 1 && dataDay <= 31 {
		result = 1
	} else {
		result = 0
	}
	return result
}

func CheckDayLimit30(dataDay int) int {
	var result int
	if dataDay >= 1 && dataDay <= 30 {
		result = 1
	} else {
		result = 0
	}
	return result
}
