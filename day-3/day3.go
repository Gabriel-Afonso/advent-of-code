package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("day3.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	lineSymbols := make(map[rune]int8)
	var result int32 = 0
	var result2 int32 = 0

	// part 1
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i])/2; j++ {
			lineSymbols[[]rune(lines[i])[j]] = 1
		}
		for j := len(lines[i]) / 2; j < len(lines[i]); j++ {
			symbol := []rune(lines[i])[j]
			if lineSymbols[symbol] == 1 {
				result += prioriry(symbol)
				break
			}
		}
		lineSymbols = make(map[rune]int8)
	}

	// part 2
	for i := 0; i < len(lines); {
		for group := 0; group < 3; i, group = i+1, group+1 {
			for j := 0; j < len(lines[i]); j++ {
				symbol := []rune(lines[i])[j]
				lineSymbols[symbol] |= 1 << group
				if group > 1 && lineSymbols[symbol] == 7 {
					result2 += prioriry(symbol)
					break
				}
			}
		}
		lineSymbols = make(map[rune]int8)
	}

	fmt.Println(result)
	fmt.Println(result2)
}

func prioriry(symbol int32) int32 {
	if symbol >= 'a' {
		return symbol - 'a' + 1
	}
	return symbol - 'A' + 27
}
