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

	fmt.Println(solve(dat, 4))
	fmt.Println(solve(dat, 14))
}

func solve(dat []byte, distinctSymbolCount int) int {
	contains := make(map[byte]int, 26)
	startI, endI := 0, 0
	for ; endI-startI != distinctSymbolCount && startI < len(dat)-distinctSymbolCount-1; endI++ {
		lastSymbolPos := contains[dat[endI]]
		contains[dat[endI]] = endI
		if lastSymbolPos != 0 && lastSymbolPos >= startI {
			startI = lastSymbolPos + 1
			continue
		}
	}
	return endI
}
