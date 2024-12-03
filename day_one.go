package main

import (
	"os"
	"strconv"
)

func dayOneResult() int {
	return totalDistance(parseFile())
}

func parseFile() ([]int, []int) {
	file, err := os.Open("resources/day_one_input")
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	line := ""

	var leftList []int
	var rightList []int
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			for _, b := range buffer[:n] {
				if b == '\n' {
					d1, d2 := parseLine(line)
					leftList = append(leftList, d1)
					rightList = append(rightList, d2)
					line = ""
				} else {
					line += string(b)
				}
			}
		}

		if err != nil {
			if line != "" {
				d1, d2 := parseLine(line)
				leftList = append(leftList, d1)
				rightList = append(rightList, d2)
			}
			break
		}
	}
	return leftList, rightList
}

func parseLine(line string) (int, int) {
	num1, num2 := "", ""
	onSecond := false
	for _, char := range line {
		if char == ' ' {
			onSecond = true
			continue
		}

		if onSecond {
			num2 += string(char)
		} else {
			num1 += string(char)
		}
	}

	d1, _ := strconv.Atoi(num1)
	d2, _ := strconv.Atoi(num2)
	return d1, d2
}

func totalDistance(leftList, rightList []int) int {
	bubbleSort(leftList)
	bubbleSort(rightList)

	total := 0
	for i := 0; i < len(leftList); i++ {
		var delta int
		if leftList[i] < rightList[i] {
			delta = rightList[i] - leftList[i]
		} else {
			delta = leftList[i] - rightList[i]
		}
		total += delta
	}
	return total
}

func bubbleSort(l []int) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l)-i-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
