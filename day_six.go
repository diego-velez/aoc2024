package main

type Coord struct {
	Row    int
	Column int
}

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

func daySixResult() (int, int) {
	var path [][]rune
	ParseFile("./resources/day_six_input", func(line string) {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		path = append(path, row)
	})

	distinctPositions, pathVisualization := distinctPositions(path, nil)
	return distinctPositions, calcPossibleObstructions(pathVisualization)
}

func calcPossibleObstructions(path [][]rune) int {
	return 0
}

func getPossibleObstructionPlaces(path [][]rune) []Coord {
	// TODO: should probably use a set for the possibleelocations
	for _, row := range path {
		for _, char := range row {
			if char == '#' {
			}
		}
	}
	return []Coord{}
}

func getPossibleObstructionsForObsticle(path [][]rune, rowIndex, columnIndex int) []Coord {
	var possibleLocations []Coord

	for thisRowIndex, row := range path {
		for thisColumnIndex, char := range row {
			// Check quadrant 1, row + 1 coming from top
			// Check quadrant 2, column + 1 coming from left
			// check quadrant 3, row - 1 coming from bottom
			// Check quadrant 4, column - 1 coming from right
			quadrant1 := thisRowIndex == rowIndex+1 && char == '|'
			quadrant2 := thisColumnIndex == columnIndex+1 && char == '-'
			quadrant3 := thisRowIndex == rowIndex-1 && char == '|'
			quadrant4 := thisColumnIndex == columnIndex-1 && char == '-'
			if quadrant1 || quadrant2 || quadrant3 || quadrant4 {
				possibleLocations = append(possibleLocations, Coord{Row: thisRowIndex, Column: thisColumnIndex})
			}
		}
	}

	return possibleLocations
}

func distinctPositions(path [][]rune, pathVis [][]rune) (int, [][]rune) {
	// NOTE: You can visualize here
	// printMap(pathVis)
	if pathVis == nil {
		pathVis = deepCopyMap(path)
	}

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
			return calcDistinctPositions(path), pathVis
		}

		blocked = path[frontRowIndex][frontColumnIndex] == '#'
		if !blocked {
			path[frontRowIndex][frontColumnIndex] = path[rowIndex][columnIndex]

			path[rowIndex][columnIndex] = 'X'

			if direction == UP || direction == DOWN {
				pathVis[frontRowIndex][frontColumnIndex] = '|'
			} else {
				pathVis[frontRowIndex][frontColumnIndex] = '-'
			}

			rowIndex, columnIndex = frontRowIndex, frontColumnIndex
		}
	}

	path[rowIndex][columnIndex] = nextDirectionChar
	pathVis[rowIndex][columnIndex] = '+'

	return distinctPositions(path, pathVis)
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

func deepCopyMap(path [][]rune) [][]rune {
	pathCopy := make([][]rune, len(path))

	for i, row := range path {
		pathCopy[i] = make([]rune, len(row))
		copy(pathCopy[i], row)
	}

	return pathCopy
}
