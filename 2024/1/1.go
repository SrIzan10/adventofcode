package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func p1(data []string) int {
	var left, right []int
	var totalDistance int

	for i := 0; i < len(data); i++ {
		splitData := strings.Split(data[i], "   ")
		splitInt, err := sliceAtoi(splitData)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		
		left = append(left, splitInt[0])
		right = append(right, splitInt[1])
	}

	// sort
	sort.Ints(left[:])
	sort.Ints(right[:])

	// continue looping by the data length
	for i := 0; i < len(data); i++ {
		modLeft, smallestLeft := remove(left, 0)
		modRight, smallestRight := remove(right, 0)

		left = modLeft
		right = modRight

		// idk why it has to use float64 with such a probably small number but i dont care
		totalDistance += int(math.Abs(float64(smallestLeft - smallestRight)))
	}

	return totalDistance
}

func p2(data []string) int {
	var left, right []int
	var totalSimilarity int

	for i := 0; i < len(data); i++ {
		splitData := strings.Split(data[i], "   ")
		splitInt, err := sliceAtoi(splitData)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		
		left = append(left, splitInt[0])
		right = append(right, splitInt[1])
	}

	// sort
	sort.Ints(left[:])
	sort.Ints(right[:])

	// here comes the funny
	occurence := frequencyMap(right)
	for i := 0; i < len(left); i++ {
		currentOcc := left[i]
		countOcc := occurence[currentOcc]

		totalSimilarity += currentOcc * countOcc
	}

	return totalSimilarity
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

// https://stackoverflow.com/a/24973180
func sliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}

// modified from https://stackoverflow.com/a/37335777
func remove(slice []int, s int) ([]int, int) {
    removedElement := slice[s]
    newSlice := append(slice[:s], slice[s+1:]...)
    return newSlice, removedElement
}

// https://www.geeksforgeeks.org/golang-program-to-find-the-frequency-of-each-element-in-an-array/
func frequencyMap(arr []int) map[int]int {
    freq := make(map[int]int)
    for _ , num :=  range arr {
        freq[num] = freq[num]+1
    }
    return freq
}