package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("day2.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	// part 1
	symbols := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	// part 2
	symbols2 := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	result := 0
	result2 := 0
	for i := 0; i < len(lines); i++ {
		result += symbols[lines[i]]
		result2 += symbols2[lines[i]]
	}
	fmt.Println(result)
	fmt.Println(result2)
}
