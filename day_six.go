package main

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

func daySixResult() int {
	var path [][]rune
	ParseFile("./resources/day_six_input", func(line string) {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		path = append(path, row)
	})
	return distinctPositions(path)
}

func distinctPositions(path [][]rune) int {
	// NOTE: You can visualize here
	// printMap(path)
	rowIndex, columnIndex := currentPosition(path)
	direction, nextDirectionChar := heading(path[rowIndex][columnIndex])

	blocked := false
	for !blocked {
		frontRowIndex, frontColumnIndex := rowIndex, columnIndex
		switch direction {
		case LEFT:
			frontColumnIndex--
		case RIGHT:
			frontColumnIndex++
		case UP:
			frontRowIndex--
		case DOWN:
			frontRowIndex++
		}

		outOfBounds := frontRowIndex < 0 || frontRowIndex >= len(path) || frontColumnIndex < 0 || frontColumnIndex >= len(path[0])
		if outOfBounds {
			return calcDistinctPositions(path)
		}

		blocked = path[frontRowIndex][frontColumnIndex] == '#'
		if !blocked {
			path[frontRowIndex][frontColumnIndex] = path[rowIndex][columnIndex]
			path[rowIndex][columnIndex] = 'X'
			rowIndex, columnIndex = frontRowIndex, frontColumnIndex
		}
	}

	path[rowIndex][columnIndex] = nextDirectionChar

	return distinctPositions(path)
}

func calcDistinctPositions(path [][]rune) int {
	total := 1
	for _, row := range path {
		for _, char := range row {
			if char == 'X' {
				total++
			}
		}
	}
	return total
}

func currentPosition(path [][]rune) (int, int) {
	for rowIndex, row := range path {
		for columnIndex, char := range row {
			if char == '<' || char == '>' || char == '^' || char == 'v' {
				return rowIndex, columnIndex
			}
		}
	}
	panic("Could not find the guard's position")
}

func heading(char rune) (Direction, rune) {
	switch char {
	case '^':
		return UP, '>'
	case '>':
		return RIGHT, 'v'
	case '<':
		return LEFT, '^'
	case 'v':
		return DOWN, '<'
	}
	panic("Guard is not a valid character")
}

func printMap(path [][]rune) {
	for _, row := range path {
		for _, char := range row {
			print(string(char))
		}
		println()
	}
	println()
}
