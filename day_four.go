package main

import (
	"strings"
)

func dayFourResult() int {
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

func occurences(data [][]string) int {
	result := 0

	for rowIndex, row := range data {
		for columnIndex, char := range row {
			if char != "X" {
				continue
			}

			result += checkHorizontal(row, columnIndex) + checkVertical(data, rowIndex, columnIndex) + checkDiagonal(data, rowIndex, columnIndex)
		}
	}

	return result
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

	canGoTop := row-3 >= 0
	canGoBottom := row+3 <= len(data)-1

	canGoLeft := column-3 >= 0
	canGoRight := column+3 <= len(data[row])-1

	// Top-right
	if canGoTop && canGoRight {
		word := "X" + data[row-1][column+1] + data[row-2][column+2] + data[row-3][column+3]
		if word == "XMAS" {
			result++
		}
	}

	// Top-left
	if canGoTop && canGoLeft {
		word := "X" + data[row-1][column-1] + data[row-2][column-2] + data[row-3][column-3]
		if word == "XMAS" {
			result++
		}
	}

	// Bottom-right
	if canGoBottom && canGoRight {
		word := "X" + data[row+1][column+1] + data[row+2][column+2] + data[row+3][column+3]
		if word == "XMAS" {
			result++
		}
	}

	// Bottom-left
	if canGoBottom && canGoLeft {
		word := "X" + data[row+1][column-1] + data[row+2][column-2] + data[row+3][column-3]
		if word == "XMAS" {
			result++
		}
	}

	return result
}
