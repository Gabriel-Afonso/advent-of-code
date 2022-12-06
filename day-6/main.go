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

	for _, distinct := range [2]int{4, 14} {
		for cursor := 0; cursor < len(dat)-distinct-1; {
			next := checkSection(dat, cursor, distinct)
			cursor += next
			if next == 0 {
				fmt.Println(cursor + distinct)
				break
			}
		}
	}
}

func checkSection(dat []byte, c int, sectionBits int) int {
	contains := make(map[byte]int, sectionBits)
	for i := 0; i < sectionBits; i++ {
		if contains[dat[c+i]] != 0 {
			return contains[dat[c+i]]
		}
		contains[dat[c+i]] = i + 1
	}
	return 0
}
