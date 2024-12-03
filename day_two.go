package main

import (
	"strconv"
)

func dayTwoResult() int {
	var reports [][]int
	ParseFile("resources/day_two_input", func(line string) {
		var report []int

		num := ""
		for _, char := range line {
			if char == ' ' {
				n, _ := strconv.Atoi(num)
				num = ""
				report = append(report, n)
				continue
			}
			num += string(char)
		}
		n, _ := strconv.Atoi(num)
		report = append(report, n)
		reports = append(reports, report)
	})
	return totalSafeReports(reports)
}

func totalSafeReports(reports [][]int) int {
	result := 0
	for _, report := range reports {
		isSafe := true
		isIncreasing := true
		for i := 0; i < len(report)-1; i++ {
			this, next := report[i], report[i+1]
			delta := delta(this, next)

			// Check delta
			if delta > 3 || delta < 1 {
				isSafe = false
				break
			}

			// Check all increasing or decreasing
			if i == 0 {
				isIncreasing = next > this
				continue
			}

			followsIncOrDecRule := (isIncreasing && next >= this) || (!isIncreasing && next <= this)
			if !followsIncOrDecRule {
				isSafe = false
				break
			}
		}

		if isSafe {
			result++
		}
	}
	return result
}

func delta(a, b int) int {
	result := a - b
	if result < 0 {
		result *= -1
	}
	return result
}
