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

	copyPath := deepCopyPath(path)
	traversedPath := traversePath(copyPath, func(_ Coord, _ Coord, _ Direction) bool { return true }, true)
	return calcDistinctPositions(traversedPath), calcPossibleObstructions(path, traversedPath)
}

func calcPossibleObstructions(ogPath, path [][]rune) int {
	p := getPossibleObstructionPlaces(path)

	result := 0
	for _, coord := range p {
		ogPathCopy := deepCopyPath(ogPath)
		result += checkCycleWithObstacle(ogPathCopy, coord)
	}

	return result
}

// TODO: incomplete
// Should return 1 if it creates a cycle, 0 otherwise
// Probably should create a function that just traverses the guard's path and takes in another
// function that runs on each move, then I could chain the functions for part 1 and 2 respectively
func checkCycleWithObstacle(path [][]rune, coord Coord) int {
	startingPosition := currentPosition(path)
	startingDirection, _ := heading(path[startingPosition.Row][startingPosition.Column])
	path[coord.Row][coord.Column] = '#'

	doesCycle, isStarting := false, true
	traversePath(path, func(currentPosition, nextPosition Coord, currentDirection Direction) bool {
		if isStarting {
			isStarting = false
			return true
		}

		// BUG: No siempre loops por el starting position pero it does loop
		if startingPosition == currentPosition && startingDirection == currentDirection {
			doesCycle = true
		}
		return !doesCycle
	}, true)

	if doesCycle {
		printPath(path)
		return 1
	}
	return 0
}

func getPossibleObstructionPlaces(path [][]rune) []Coord {
	var possibleObstructionPlaces map[Coord]bool = make(map[Coord]bool)
	for rowIndex, row := range path {
		for columnIndex, char := range row {
			if char == '#' {
				p := getPossibleObstructionsForObsticle(path, rowIndex, columnIndex)
				addToSet(&possibleObstructionPlaces, p)
			}
		}
	}

	var result []Coord
	for coord := range possibleObstructionPlaces {
		result = append(result, coord)
	}
	return result
}

func addToSet(out *map[Coord]bool, coords []Coord) {
	for _, coord := range coords {
		if !(*out)[coord] {
			(*out)[coord] = true
		}
	}
}

func getPossibleObstructionsForObsticle(path [][]rune, rowIndex, columnIndex int) []Coord {
	var possibleLocations []Coord

	for thisRowIndex, row := range path {
		for thisColumnIndex, char := range row {
			// Check quadrant 1, row + 1 coming from top
			// Check quadrant 2, column + 1 coming from left
			// check quadrant 3, row - 1 coming from bottom
			// Check quadrant 4, column - 1 coming from right
			quadrant1 := thisRowIndex < len(path) && thisRowIndex == rowIndex+1 && char == '|'
			quadrant2 := thisColumnIndex < len(row) && thisColumnIndex == columnIndex+1 && char == '-'
			quadrant3 := thisRowIndex >= 0 && thisRowIndex == rowIndex-1 && char == '|'
			quadrant4 := thisColumnIndex >= 0 && thisColumnIndex == columnIndex-1 && char == '-'
			if quadrant1 || quadrant2 || quadrant3 || quadrant4 {
				possibleLocations = append(possibleLocations, Coord{Row: thisRowIndex, Column: thisColumnIndex})
			}
		}
	}

	return possibleLocations
}

func traversePath(path [][]rune, shouldContinue func(Coord, Coord, Direction) bool, starting bool) [][]rune {
	// NOTE: You can visualize here
	// printPath(path)
	currentPosition := currentPosition(path)
	currentDirection, nextChar := heading(path[currentPosition.Row][currentPosition.Column])

	blocked := false
	firstTime := true
	for !blocked {
		nextPosition := currentPosition
		switch currentDirection {
		case LEFT:
			nextPosition.Column--
		case RIGHT:
			nextPosition.Column++
		case UP:
			nextPosition.Row--
		case DOWN:
			nextPosition.Row++
		}

		if !shouldContinue(currentPosition, nextPosition, currentDirection) {
			return path
		}

		outOfBounds := nextPosition.Row < 0 || nextPosition.Row >= len(path) || nextPosition.Column < 0 || nextPosition.Column >= len(path[0])
		if outOfBounds {
			return path
		}

		blocked = path[nextPosition.Row][nextPosition.Column] == '#'
		if !blocked {
			path[nextPosition.Row][nextPosition.Column] = path[currentPosition.Row][currentPosition.Column]

			if firstTime && !starting {
				firstTime = false
				starting = false
				path[currentPosition.Row][currentPosition.Column] = '+'
			} else if currentDirection == UP || currentDirection == DOWN {
				path[currentPosition.Row][currentPosition.Column] = '|'
			} else {
				path[currentPosition.Row][currentPosition.Column] = '-'
			}

			currentPosition = nextPosition
		}
	}

	path[currentPosition.Row][currentPosition.Column] = nextChar

	return traversePath(path, shouldContinue, false)
}

func calcDistinctPositions(path [][]rune) int {
	total := 1
	for _, row := range path {
		for _, char := range row {
			if char == '|' || char == '-' || char == '+' {
				total++
			}
		}
	}
	return total
}

func currentPosition(path [][]rune) Coord {
	for rowIndex, row := range path {
		for columnIndex, char := range row {
			if char == '<' || char == '>' || char == '^' || char == 'v' {
				return Coord{Row: rowIndex, Column: columnIndex}
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

func printPath(path [][]rune) {
	for _, row := range path {
		for _, char := range row {
			print(string(char))
		}
		println()
	}
	println()
}

func deepCopyPath(path [][]rune) [][]rune {
	pathCopy := make([][]rune, len(path))

	for i, row := range path {
		pathCopy[i] = make([]rune, len(row))
		copy(pathCopy[i], row)
	}

	return pathCopy
}
