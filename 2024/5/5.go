package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func p1(data []string) int {
	
}

/* func p2(data []string) int {
	
} */

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1:", p1(lines))
	// fmt.Println("Part 2:", p2(lines))
}

func getRulesAndUpdates(data []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := make([][]int, 0)
	for i := range updates {
		updates[i] = make([]int, 0)
	}

	isParsingRules := true
	for _, line := range data {
		if line == "" {
			isParsingRules = false
			continue
		}

		if isParsingRules {
			splitted := strings.Split(line, "|")

			left, _ := strconv.Atoi(splitted[0])
			right, _ := strconv.Atoi(splitted[1])

			rules[left] = append(rules[left], right)
		} else {
			splits := strings.Split(line, ",")
			update := make([]int, 0, len(splits))

			for _, i := range splits {
				n, _ := strconv.Atoi(i)

				update = append(update, n)
			}

			updates = append(updates, update)
		}
	}
	return rules, updates
}

func calc(rules map[int][]int, updates [][]int) {
	
}