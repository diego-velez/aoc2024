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

func dayFiveResult() int {
	var rules []Rule
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

func doIt(rules []Rule, pages [][]int) int {
	// Basicamente puedo buscar por cada page de atras pa lante, despues buscando ese page rule
	// usando ese como el before, salvo todos los after en same array, despues voy chequiando
	// si esta en array, si esta pues no follows the rules pq vino antes pero se supone que viniera
	// despues
	total := 0

	for _, instructions := range pages {
		var allAfter []int
		validInstruction := true

		for i := len(instructions) - 1; i >= 0; i-- {
			page := instructions[i]
			if slices.Contains(allAfter, page) {
				validInstruction = false
				break
			}
			getAfters(page, rules, &allAfter)
		}

		if !validInstruction {
			continue
		}
		midIndex := (len(instructions) - 1) / 2
		total += instructions[midIndex]
	}

	return total
}

func getAfters(page int, rules []Rule, out *[]int) {
	for _, rule := range rules {
		if rule.Before == page {
			*out = append(*out, rule.After)
		}
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

func parseRules(line string, out *[]Rule) {
	parseLine := strings.Split(line, "|")
	num1, _ := strconv.Atoi(parseLine[0])
	num2, _ := strconv.Atoi(parseLine[1])
	*out = append(*out, Rule{Before: num1, After: num2})
}
