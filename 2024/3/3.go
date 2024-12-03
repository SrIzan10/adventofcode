package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func p1(data string) int {
	mult, _ := matchMul(data)
	sum := 0
	for _, value := range mult {
		sum += value
	}
	return sum
}

func p2(data string) int {
	sum := 0
	mulMult, mulIndexes := matchMul(data)

	matchDos, _ := regexp.Compile(`do+\(\)`)
	matchDonts, _ := regexp.Compile(`don't+\(\)`)
	doIndex := matchDos.FindAllStringIndex(data, -1)
	dontIndex := matchDonts.FindAllStringIndex(data, -1)
	resolvedDos := sliceLast2dValue(doIndex)
	resolvedDonts := sliceLast2dValue(dontIndex)
	resolvedMul := sliceLast2dValue(mulIndexes)

	for i, mulIndex := range resolvedMul {
        lastControl := -1
        lastWasDo := true

        // find most recent control statement
        for _, doIdx := range resolvedDos {
            if doIdx < mulIndex && doIdx > lastControl {
                lastControl = doIdx
                lastWasDo = true
            }
        }
        for _, dontIdx := range resolvedDonts {
            if dontIdx < mulIndex && dontIdx > lastControl {
                lastControl = dontIdx
                lastWasDo = false
            }
        }

        if lastWasDo {
            sum += mulMult[i]
        }
	}

	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Part 1:", p1(string(data)))
	fmt.Println("Part 2:", p2(string(data)))
}

func matchMul(data string) (mults []int, index [][]int) {
	var multiplications []int
	match, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	resolved := match.FindAllString(data, -1)
	resolvedIndex := match.FindAllStringIndex(data, -1)

	numRegex, _ := regexp.Compile(`\d+,\d+`)
	for _, resolves := range resolved {
		num := strings.Split(numRegex.FindAllString(resolves, -1)[0], ",")
		n0, _ := strconv.Atoi(num[0])
		n1, _ := strconv.Atoi(num[1])
		multiplications = append(multiplications, n0 * n1)
	}
	return multiplications, resolvedIndex
}

func sliceLast2dValue(array [][]int) []int {
	var finalList []int
	for _, value1 := range array {
		finalList = append(finalList, value1[1])
	}
	return finalList
}