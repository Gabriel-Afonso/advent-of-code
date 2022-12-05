package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	// part 1
	result := 0
	buffer := 0
	for i := 0; i < len(lines); i++ {
		intVar, err := strconv.Atoi(lines[i])
		if err == nil {
			buffer += intVar
			continue
		}
		if buffer > result {
			result = buffer
		}
		buffer = 0
	}

	// part 2
	var resultList []int
	buffer = 0
	for i := 0; i < len(lines); i++ {
		intVar, err := strconv.Atoi(lines[i])
		if err == nil {
			buffer += intVar
			continue
		}
		resultList = append(resultList, buffer)
		buffer = 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(resultList)))

	fmt.Println(result)
	fmt.Println(resultList[0] + resultList[1] + resultList[2])
}
