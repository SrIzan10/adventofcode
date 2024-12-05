package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func p1(data []string, directions [][2]int) int {
	numRows := len(data)
	numCols := len(data[0])
	totalOccurrences := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			for _, direction := range directions {
				if hasXmas(row, col, direction, data) {
					totalOccurrences++
				}
			}
		}
	}
	return totalOccurrences
}

func p2(data []string) int {
    count := 0
    rows := len(data)
    cols := len(data[0])

    for row := 1; row < rows-1; row++ {
        for col := 1; col < cols-1; col++ {
            if data[row][col] != 'A' {
                continue
            }

            // check all four diagonal combinations
            topLeft := string([]byte{data[row-1][col-1], data[row][col], data[row+1][col+1]})
            topRight := string([]byte{data[row-1][col+1], data[row][col], data[row+1][col-1]})

            if isValidMAS(topLeft) && isValidMAS(topRight) {
                count++
            }
        }
    }
    return count
}

func main() {
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1:", p1(lines, directions))
	fmt.Println("Part 2:", p2(lines))
}

func hasXmas(startRow int, startCol int, direction [2]int, data []string) bool {
	deltaRow, deltaCol := direction[0], direction[1]
	for index, char := range "XMAS" {
		currentRow := startRow + index*deltaRow
		currentCol := startCol + index*deltaCol
		numRows := len(data)
		numCols := len(data[0])

		if currentRow < 0 || currentRow >= numRows || currentCol < 0 || currentCol >= numCols {
			return false
		}
		if data[currentRow][currentCol] != byte(char) {
			return false
		}
	}
	return true
}

func sliceLast2dValue(array [][]int) []int {
	var finalList []int
	for _, value1 := range array {
		finalList = append(finalList, value1[1])
	}
	return finalList
}

func checkCharacter(data []string, rowIndex, colIndex, rowOffset, colOffset int, target []string) bool {
    newRow := rowIndex + rowOffset
    newCol := colIndex + colOffset
    if newRow >= 0 && newRow < len(data) && newCol >= 0 && newCol < len(data[0]) {
        return slices.Contains(target, string(data[newRow][newCol]))
    }
    return false
}

func isValidMAS(s string) bool {
    return s == "MAS" || s == "SAM"
}