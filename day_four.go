package main

import (
	"strings"
)

func dayFourResult() (int, int) {
	var data [][]string
	ParseFile("resources/day_four_input", func(line string) {
		var row []string

		for _, char := range line {
			row = append(row, string(char))
		}

		data = append(data, row)
	})
	return occurences(data)
}

func occurences(data [][]string) (int, int) {
	resultPartOne := 0
	resultPartTwo := 0

	for rowIndex, row := range data {
		for columnIndex, char := range row {
			// We look for the X of XMAS for the first part, and the center A for X-MAS for the second part
			if char == "X" {
				resultPartOne += checkHorizontal(row, columnIndex) + checkVertical(data, rowIndex, columnIndex) + checkDiagonal(data, rowIndex, columnIndex)
			} else if char == "A" {
				resultPartTwo += checkCrossMas(data, rowIndex, columnIndex)
			}
		}
	}

	return resultPartOne, resultPartTwo
}

func checkHorizontal(row []string, i int) int {
	result := 0

	canGoLeft := i-3 >= 0
	if canGoLeft {
		left := row[i-3 : i+1]
		leftWord := strings.Join(left, "")
		if leftWord == "SAMX" {
			result++
		}
	}

	canGoRight := i+3 <= len(row)-1
	if canGoRight {
		right := row[i : i+4]
		rightWord := strings.Join(right, "")
		if rightWord == "XMAS" {
			result++
		}
	}

	return result
}

func checkVertical(data [][]string, row, column int) int {
	result := 0

	canGoUp := row-3 >= 0
	if canGoUp {
		top := ""
		for i := row; i >= row-3; i-- {
			top += data[i][column]
		}
		if top == "XMAS" {
			result++
		}
	}

	canGoDown := row+3 <= len(data)-1
	if canGoDown {
		bottom := ""
		for i := row; i <= row+3; i++ {
			bottom += data[i][column]
		}
		if bottom == "XMAS" {
			result++
		}
	}

	return result
}

func checkDiagonal(data [][]string, row, column int) int {
	result := 0

	canGoUp := row-3 >= 0
	canGoDown := row+3 <= len(data)-1

	canGoLeft := column-3 >= 0
	canGoRight := column+3 <= len(data[row])-1

	// Top-right
	if canGoUp && canGoRight {
		word := "X" + data[row-1][column+1] + data[row-2][column+2] + data[row-3][column+3]
		if word == "XMAS" {
			result++
		}
	}

	// Top-left
	if canGoUp && canGoLeft {
		word := "X" + data[row-1][column-1] + data[row-2][column-2] + data[row-3][column-3]
		if word == "XMAS" {
			result++
		}
	}

	// Bottom-right
	if canGoDown && canGoRight {
		word := "X" + data[row+1][column+1] + data[row+2][column+2] + data[row+3][column+3]
		if word == "XMAS" {
			result++
		}
	}

	// Bottom-left
	if canGoDown && canGoLeft {
		word := "X" + data[row+1][column-1] + data[row+2][column-2] + data[row+3][column-3]
		if word == "XMAS" {
			result++
		}
	}

	return result
}

func checkCrossMas(data [][]string, row, column int) int {
	canGoLeft := column-1 >= 0
	canGoRight := column+1 <= len(data[row])-1

	canGoUp := row-1 >= 0
	canGoDown := row+1 <= len(data)-1

	if !canGoLeft || !canGoRight || !canGoUp || !canGoDown {
		return 0
	}

	topLeft := data[row-1][column-1]
	topRight := data[row-1][column+1]

	bottomLeft := data[row+1][column-1]
	bottomRight := data[row+1][column+1]

	// Goes to the right
	crossMasLeft := topLeft == "S" && bottomLeft == "S" && topRight == "M" && bottomRight == "M"
	crossMasRight := topLeft == "M" && bottomLeft == "M" && topRight == "S" && bottomRight == "S"
	crossMasUp := bottomLeft == "M" && bottomRight == "M" && topLeft == "S" && topRight == "S"
	crossMasDown := bottomLeft == "S" && bottomRight == "S" && topLeft == "M" && topRight == "M"

	if crossMasLeft || crossMasRight || crossMasUp || crossMasDown {
		return 1
	}

	return 0
}
