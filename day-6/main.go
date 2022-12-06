package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("day6.txt")

	if err != nil {
		panic(err)
	}

	// part 1
	distinctCount := 4
	for cursor := 0; cursor < len(dat)-3; {
		next := checkSection(dat, cursor, distinctCount)
		cursor += next
		if next == 0 {
			fmt.Println(cursor + distinctCount)
			break
		}
	}

	// part 2
	distinctCount = 14
	for cursor := 0; cursor < len(dat)-3; {
		next := checkSection(dat, cursor, distinctCount)
		cursor += next
		if next == 0 {
			fmt.Println(cursor + distinctCount)
			break
		}
	}
}

func checkSection(dat []byte, c int, sectionBits int) int {
	for i := 0; i < sectionBits-1; i++ {
		for j := i + 1; j < sectionBits; j++ {
			if dat[c+i] == dat[c+j] {
				return i + 1
			}
		}
	}
	return 0
}
