package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("day4.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	result1 := 0
	result2 := 0
	for i := 0; i < len(lines); i++ {
		section := strings.Split(lines[i], ",")

		elf1Section := strings.Split(section[0], "-")
		elf1Start, _ := strconv.Atoi(elf1Section[0])
		elf1End, _ := strconv.Atoi(elf1Section[1])

		elf2Section := strings.Split(section[1], "-")
		elf2Start, _ := strconv.Atoi(elf2Section[0])
		elf2End, _ := strconv.Atoi(elf2Section[1])

		// part 1
		if elf1Start <= elf2Start && elf1End >= elf2End || elf2Start <= elf1Start && elf2End >= elf1End {
			result1++
		}

		// part 2
		if elf1End >= elf2Start && elf2End >= elf1Start {
			result2++
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
