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
	outer:
		for c := 0; c < len(dat)-distinct-1; {
			contains := make(map[byte]int, distinct)
			for i := 0; i < distinct; i++ {
				if contains[dat[c+i]] != 0 {
					c += contains[dat[c+i]]
					continue outer
				}
				contains[dat[c+i]] = i + 1
			}
			fmt.Println(c + distinct)
			break
		}
	}
}
