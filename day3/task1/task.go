package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day3/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var grid [][]rune

	for _, row := range rows {
		grid = append(grid, []rune(row))
	}

	treeCount := 0

	x := 0
	y := 0

	for y < len(grid) {
		if grid[y][x%len(grid[0])] == '#' {
			treeCount++
		}

		x += 3
		y++
	}

	fmt.Printf("Solution is: %d\n", treeCount)
}
