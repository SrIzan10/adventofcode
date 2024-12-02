package main

import (
	"fmt"
	"os"
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