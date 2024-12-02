package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// each line is one report, one number in a report is a level
func p1(data []string) int {
	safeReports := 0
    for _, line := range data {
        levels := sliceAtoi(strings.Fields(line))
        if isValidSequence(levels) {
            safeReports++
        }
    }
    return safeReports
}

func p2(data []string) int {
    safeReports := 0
    
    for _, line := range data {
        levels := sliceAtoi(strings.Fields(line))
        
        // check original sequence
        if isValidSequence(levels) {
            safeReports++
            continue
        }
        
        // try removing one level at a time
        for i := range levels {
            newLevels := make([]int, 0, len(levels)-1)
            newLevels = append(newLevels, levels[:i]...)
            newLevels = append(newLevels, levels[i+1:]...)
            
            if isValidSequence(newLevels) {
                safeReports++
                break
            }
        }
    }
    return safeReports
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1:", p1(lines))
	fmt.Println("Part 2:", p2(lines))
}

// HELPER FUNCTIONS

// day1, modified to ignore error https://stackoverflow.com/a/24973180
func sliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, _ := strconv.Atoi(a)
		si = append(si, i)
	}
	return si
}

func isValidSequence(levels []int) bool {
    if len(levels) < 2 {
        return false
    }
    
    isIncreasing := levels[1] > levels[0]
    
    for i := 1; i < len(levels); i++ {
        diff := levels[i] - levels[i-1]
        if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
            return false
        }
        if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
            return false
        }
    }
    return true
}