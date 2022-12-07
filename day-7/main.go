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
	parent *Node
	name   string
	dirs   map[string]*Node
	size   int
}

func main() {
	dat, err := os.ReadFile("day7.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	rootNode := &Node{
		parent: nil,
		name:   "/",
	}
	currentDir := rootNode
	for _, line := range lines[2:] {
		switch line[:4] {
		case "$ ls":
			continue
		case "$ cd":
			if len(line) > 6 && line[5:7] == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.dirs[line[5:]]
			}
		case "dir ":
			if currentDir.dirs == nil {
				currentDir.dirs = make(map[string]*Node)
			}
			currentDir.dirs[line[4:]] = &Node{
				parent: currentDir,
				name:   line[4:],
			}
		default:
			split := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(split[0])
			currentDir.size += fileSize
		}
	}
	calcDirSize(rootNode)

	fmt.Println(part1(rootNode, atMost))
	fmt.Println(part2(rootNode, (space-rootNode.size-fit)*-1))
	print(rootNode, 0)
}

func calcDirSize(node *Node) {
	for _, child := range node.dirs {
		calcDirSize(child)
		node.size += child.size
	}
}

func print(node *Node, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s: ", node.name)
	fmt.Println(node.size)
	for _, node := range node.dirs {
		print(node, indent+2)
	}
}

func part2(node *Node, fit int) int {
	bucket := node.size
	findFit(node, fit, &bucket)
	return bucket
}

func findFit(node *Node, fit int, bucket *int) {
	for _, node := range node.dirs {
		findFit(node, fit, bucket)
	}
	if node.size > fit && node.size < *bucket {
		*bucket = node.size
	}
}

func part1(node *Node, atMost int) int {
	bucket := 0
	traversAndCollect(node, atMost, &bucket)
	return bucket
}

func traversAndCollect(node *Node, atMost int, bucket *int) {
	for _, node := range node.dirs {
		traversAndCollect(node, atMost, bucket)
	}
	if node.size < atMost {
		*bucket += node.size
	}
}
