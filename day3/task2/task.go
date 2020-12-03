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

	var treeCounts []int
	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	for _, slope := range slopes {
		treeCount := 0

		x := 0
		y := 0

		for y < len(grid) {
			if grid[y][x%len(grid[0])] == '#' {
				treeCount++
			}

			x += slope[0]
			y += slope[1]
		}
		treeCounts = append(treeCounts, treeCount)
	}

	sol := 1

	for _, cnt := range treeCounts {
		sol *= cnt
	}

	fmt.Printf("Solution is: %d\n", sol)
}
