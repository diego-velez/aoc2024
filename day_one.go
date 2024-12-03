package main

import (
	"strconv"
)

func dayOneResult() (int, int) {
	var leftList, rightList []int
	ParseFile("resources/day_one_input", func(line string) {
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
		leftList = append(leftList, d1)
		rightList = append(rightList, d2)
	})

	bubbleSort(leftList)
	bubbleSort(rightList)
	return totalDistance(leftList, rightList), similarityScore(leftList, rightList)
}

func similarityScore(leftList, rightList []int) int {
	var occurences map[int]int = make(map[int]int)
	for _, num := range rightList {
		if occurences[num] == 0 {
			occurences[num] = 1
		} else {
			occurences[num]++
		}
	}

	result := 0
	for _, num := range leftList {
		result += num * occurences[num]
	}
	return result
}

func totalDistance(leftList, rightList []int) int {
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
