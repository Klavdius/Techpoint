package main

import (
	"fmt"
)

func main() {
	var amountRecord int
	fmt.Scan(&amountRecord)
	for i := 0; i < amountRecord; i++ {
		var (
			inputLine  string
			line       string
			saveLine   string
			field      = []string{}
			cursor     = 0
			numberLine = 0
		)

		fmt.Scan(&inputLine)
		lenInputLine := len(inputLine)
		for inner := 0; inner < lenInputLine; inner++ {
			sign := inputLine[:1]
			switch sign {
			case "D":
				exist := CheckingLineInField(field, numberLine, +1)
				if exist {
					field, line, saveLine = DownLine(field, line, saveLine, numberLine)
					numberLine++
				}
			case "U":
				exist := CheckingLineInField(field, numberLine, -1)
				if exist {
					field, line, saveLine = UpLine(field, line, saveLine, numberLine)
					numberLine--
				}
			case "E":
				line, saveLine = AllInLine(line, saveLine)
			case "B":
				line, saveLine = AllInSave(line, saveLine)
			case "N":
				field, line, saveLine = NewLine(field, line, saveLine, numberLine)
				numberLine++
				line, saveLine = AllInSave(line, saveLine)
			case "L":
				line, saveLine = AddSaveLine(line, saveLine)
			case "R":
				line, saveLine = DecreaseSave(line, saveLine)
			default:
				line = line + sign
				cursor++
			}
			inputLine = inputLine[1:]
		}
		line = line + saveLine
		field = CheckingField(field, numberLine)
		field[numberLine] = line

		PrintField(field)
		fmt.Println("-")
	}
}

func AddSaveLine(line string, saveLine string) (string, string) {
	if len(line) != 0 {
		sign := line[len(line)-1:]
		line = line[:len(line)-1]
		saveLine = sign + saveLine
	}

	return line, saveLine
}

func DecreaseSave(line string, save string) (string, string) {
	if len(save) != 0 {
		sign := save[:1]
		save = save[1:]
		line = line + sign
	}

	return line, save
}

func PrintField(field []string) {
	for _, v := range field {
		for _, sign := range v {
			fmt.Printf("%c", sign)
		}
		fmt.Println()
	}
}

func NewLine(field []string, line string, saveLine string, numberLine int) ([]string, string, string) {
	field = append(field, "")
	for i := len(field) - 1; i >= numberLine+1; i-- {
		field[i] = field[i-1]
		field[i-1] = ""
	}
	field[numberLine] = line
	line = saveLine
	field = CheckingField(field, numberLine+1)
	if line != "" {
		field[numberLine+1] = line
	}
	saveLine = ""
	return field, line, saveLine
}

func AllInSave(line string, saveLine string) (string, string) {
	for len(line) != 0 {
		line, saveLine = AddSaveLine(line, saveLine)
	}
	return line, saveLine
}

func AllInLine(line string, saveLine string) (string, string) {
	for len(saveLine) != 0 {
		line, saveLine = DecreaseSave(line, saveLine)
	}
	return line, saveLine
}

func UpLine(field []string, line string, saveLine string, numberLine int) ([]string, string, string) {
	cursor := len(line)
	line = line + saveLine
	field = CheckingField(field, numberLine)
	field[numberLine] = line
	line = field[numberLine-1]
	saveLine = ""
	if len(line) > cursor {
		for i := len(line); i > cursor; i-- {
			line, saveLine = AddSaveLine(line, saveLine)
		}
	}

	return field, line, saveLine
}

func DownLine(field []string, line string, saveLine string, numberLine int) ([]string, string, string) {
	cursor := len(line)
	line = line + saveLine
	field[numberLine] = line
	line = field[numberLine+1]
	saveLine = ""
	if len(line) > cursor {
		for i := len(line); i > cursor; i-- {
			line, saveLine = AddSaveLine(line, saveLine)
		}
	}
	return field, line, saveLine
}

func CheckingLineInField(field []string, numberLine int, route int) bool {
	result := false
	for i, _ := range field {
		if i == numberLine+route {
			result = true
		}
	}
	return result
}
func CheckingField(field []string, numberLine int) []string {
	panicRang := true
	for i, _ := range field {
		if i == numberLine {
			panicRang = false
		}
	}
	if panicRang {
		field = append(field, "")
	}
	return field
}