package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	input, err := os.ReadFile("day6.txt")

	if err != nil {
		panic(err)
	}

	result1, result2 := 0, 0

	start := time.Now()
	for i := 0; i < 10000; i++ {
		result1 = solve(input, 4)
	}
	elapsed1 := time.Since(start) / 10000

	start = time.Now()
	for i := 0; i < 10000; i++ {
		result2 = solve(input, 14)
	}
	elapsed2 := time.Since(start) / 10000

	fmt.Println(result1)
	fmt.Println(elapsed1)
	fmt.Println(result2)
	fmt.Println(elapsed2)
}

func solve(arr []byte, markerLen int) int {
	symbol := [26]int{}
	anchor, i := 0, 0
	for ; i-anchor != markerLen && anchor < len(arr)-markerLen-1; i++ {
		lastSymbolPos := symbol[arr[i]-97]
		if lastSymbolPos != 0 && lastSymbolPos >= anchor {
			anchor = lastSymbolPos + 1
		}
		symbol[arr[i]-97] = i
	}
	return i
}
