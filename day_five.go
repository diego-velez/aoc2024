package main

import (
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func dayFiveResult() (int, int) {
	rules := make(map[int][]int)
	var pages [][]int
	scanningPages := false
	ParseFile("resources/day_five_input", func(line string) {
		if line == "" {
			scanningPages = true
			return
		}

		if scanningPages {
			parsePages(line, &pages)
			return
		}

		parseRules(line, &rules)
	})
	return doIt(rules, pages)
}

func doIt(rules map[int][]int, pages [][]int) (int, int) {
	// Basicamente puedo buscar por cada page de atras pa lante, despues buscando ese page rule
	// usando ese como el before, salvo todos los after en same array, despues voy chequiando
	// si esta en array, si esta pues no follows the rules pq vino antes pero se supone que viniera
	// despues
	totalCorrectInstructions := 0
	totalIncorrectInstructions := 0

	for _, instructions := range pages {
		var allAfter []int
		validInstruction := true

		for i := len(instructions) - 1; i >= 0; i-- {
			page := instructions[i]
			if slices.Contains(allAfter, page) {
				fixOrder(&instructions, i, rules)
				validInstruction = false
				i = len(instructions)
				allAfter = []int{}
			} else {
				getAfters(page, rules, &allAfter)
			}
		}

		midIndex := (len(instructions) - 1) / 2
		if validInstruction {
			totalCorrectInstructions += instructions[midIndex]
		} else {
			totalIncorrectInstructions += instructions[midIndex]
		}
	}

	return totalCorrectInstructions, totalIncorrectInstructions
}

// fixOrder will swap pages that violate the rules, it will check what rule was broken and swap
// them from instructions
func fixOrder(instructions *[]int, i int, rules map[int][]int) {
	page := (*instructions)[i]

	// Get all before pages from which the problematic page is supposed to come after
	var losBefores []int
	for before, afters := range rules {
		for _, after := range afters {
			if page == after {
				losBefores = append(losBefores, before)
			}
		}
	}

	// Get the before page that comes after the i page, causing the problem
	elBeforeAfterIndex := -1
	for _, before := range losBefores {
		for j := i + 1; j < len(*instructions); j++ {
			unBeforeAfter := (*instructions)[j]
			if before == unBeforeAfter {
				elBeforeAfterIndex = j
				break
			}
		}

		if elBeforeAfterIndex != -1 {
			break
		}
	}

	// Perform the swap
	(*instructions)[i], (*instructions)[elBeforeAfterIndex] = (*instructions)[elBeforeAfterIndex], (*instructions)[i]
}

// getAfters will get all the pages that are supposed to come after page according to the rules
// and append all of them to out
func getAfters(page int, rules map[int][]int, out *[]int) {
	for _, rule := range rules[page] {
		*out = append(*out, rule)
	}
}

func parsePages(line string, out *[][]int) {
	var pages []int
	for _, page := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(page)
		pages = append(pages, num)
	}
	*out = append(*out, pages)
}

func parseRules(line string, out *map[int][]int) {
	parseLine := strings.Split(line, "|")
	num1, _ := strconv.Atoi(parseLine[0])
	num2, _ := strconv.Atoi(parseLine[1])

	if (*out)[num1] == nil {
		(*out)[num1] = []int{}
	}

	(*out)[num1] = append((*out)[num1], num2)
}
