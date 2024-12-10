package main

import (
	"strconv"
	"strings"
)

type Operation struct {
	TestValue int
	Numbers   []int
}

func daySevenResult() int {
	var operations []Operation
	ParseFile("./resources/day_seven_input", func(line string) {
		parsedLine := strings.Split(line, ": ")
		parsedNumbers := strings.Split(parsedLine[1], " ")
		testValue, _ := strconv.Atoi(parsedLine[0])

		var numbers []int
		for _, strNum := range parsedNumbers {
			num, _ := strconv.Atoi(strNum)
			numbers = append(numbers, num)
		}

		operations = append(operations, Operation{
			TestValue: testValue,
			Numbers:   numbers,
		})
	})
	return totalCalibrationResult(operations)
}

func totalCalibrationResult(operations []Operation) int {
	total := 0
	for _, operation := range operations {
		if validOperation(operation) {
			total += operation.TestValue
		}
	}
	return total
}

func validOperation(operation Operation) bool {
	return dvt(operation.TestValue, operation.Numbers[0], operation.Numbers[1:])
}

func dvt(testValue int, sumSoFar int, nums []int) bool {
	if sumSoFar == testValue {
		return true
	} else if len(nums) == 0 {
		return false
	}

	sum := sumSoFar + nums[0]
	mult := sumSoFar * nums[0]

	otherNums := nums[1:]

	return dvt(testValue, sum, otherNums) || dvt(testValue, mult, otherNums)
}
