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

	start := time.Now()
	result1, result2 := 0, 0
	for i := 0; i < 1000; i++ {
		result1 = solve(input, 4)
		result2 = solve(input, 14)
	}
	elapsed := time.Since(start) / 1000

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(elapsed)
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
