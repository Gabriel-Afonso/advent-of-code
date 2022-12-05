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
	symbols := make(map[string]int)
	symbols["A X"] = 4
	symbols["A Y"] = 8
	symbols["A Z"] = 3
	symbols["B X"] = 1
	symbols["B Y"] = 5
	symbols["B Z"] = 9
	symbols["C X"] = 7
	symbols["C Y"] = 2
	symbols["C Z"] = 6

	// part 2
	symbols2 := make(map[string]int, 9)
	symbols2["A X"] = 3
	symbols2["A Y"] = 4
	symbols2["A Z"] = 8
	symbols2["B X"] = 1
	symbols2["B Y"] = 5
	symbols2["B Z"] = 9
	symbols2["C X"] = 2
	symbols2["C Y"] = 6
	symbols2["C Z"] = 7

	result := 0
	result2 := 0
	for i := 0; i < len(lines); i++ {
		result += symbols[lines[i]]
		result2 += symbols2[lines[i]]
	}
	fmt.Println(result)
	fmt.Println(result2)
}
