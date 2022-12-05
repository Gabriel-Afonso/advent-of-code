package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	amount    int
	fromStack int
	toStack   int
}

func main() {
	dat, err := os.ReadFile("day5.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	crates := make([][]rune, 0)
	moves := make([]Move, 10)

	// parse file
	for _, line := range lines {
		if len(line) > 0 {
			switch []rune(line)[0] {
			case '[':
				// create and fill stacks
				for j := 0; j <= len(line)/4; j++ {
					symbol := []rune(line)[j*4+1]
					if len(crates) <= j {
						crates = append(crates, []rune{})
					}
					if symbol != ' ' {
						crates[j] = append(crates[j], symbol)
					}
				}
			case 'm':
				// create moves
				sections := strings.Split(line, " ")
				amount, _ := strconv.Atoi(sections[1])
				fromStack, _ := strconv.Atoi(sections[3])
				toStack, _ := strconv.Atoi(sections[5])
				moves = append(moves, Move{amount, fromStack - 1, toStack - 1})
			}
		}
	}
	// create copy of crates for second part
	crates2 := make([][]rune, len(crates))
	copy(crates2, crates)

	printResult(part1(crates, moves))
	printResult(part2(crates2, moves))
}

func printResult(crates [][]rune) {
	for _, row := range crates {
		fmt.Print(string(row[:1]))
	}
	fmt.Println()
}

func part1(crates [][]rune, moves []Move) [][]rune {
	for _, m := range moves {
		for i := 0; i < m.amount; i++ {
			// push to new stack
			crates[m.toStack] = append([]rune{crates[m.fromStack][0]}, crates[m.toStack]...)
			// pop from current stack
			crates[m.fromStack] = crates[m.fromStack][1:]
		}
	}
	return crates
}

func part2(crates [][]rune, moves []Move) [][]rune {
	for _, m := range moves {
		// push to new stack
		crates[m.toStack] = append(crates[m.fromStack][:m.amount:m.amount], crates[m.toStack]...)
		// pop from current stack
		crates[m.fromStack] = crates[m.fromStack][m.amount:]
	}
	return crates
}
