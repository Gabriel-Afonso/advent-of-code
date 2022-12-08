package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const atMost = 100000
const space = 70000000
const fit = 30000000

type Node struct {
	parent   *Node
	children []*Node
	size     int
}

func main() {
	dat, err := os.ReadFile("day7.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	rootNode := &Node{parent: nil}
	currentDir := rootNode
	for _, line := range lines[2:] {
		switch line[:4] {
		case "$ ls":
		case "dir ":
		case "$ cd":
			if line[5:7] == ".." {
				currentDir = currentDir.parent
			} else {
				newNode := &Node{parent: currentDir}
				currentDir.children = append(currentDir.children, newNode)
				currentDir = newNode
			}
		default:
			split := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(split[0])
			currentDir.size += fileSize
		}
	}
	calcDirSize(rootNode)

	fmt.Println(solve1(rootNode, atMost))
	fmt.Println(solve2(rootNode, fit-(space-rootNode.size)))
}

func calcDirSize(node *Node) {
	for _, child := range node.children {
		calcDirSize(child)
		node.size += child.size
	}
}

func solve2(node *Node, fit int) int {
	bucket := node.size
	findFit(node, fit, &bucket)
	return bucket
}

func findFit(node *Node, fit int, bucket *int) {
	for _, node := range node.children {
		findFit(node, fit, bucket)
	}
	if node.size > fit && node.size < *bucket {
		*bucket = node.size
	}
}

func solve1(node *Node, atMost int) int {
	bucket := 0
	traversAndCollect(node, atMost, &bucket)
	return bucket
}

func traversAndCollect(node *Node, atMost int, bucket *int) {
	for _, node := range node.children {
		traversAndCollect(node, atMost, bucket)
	}
	if node.size < atMost {
		*bucket += node.size
	}
}
