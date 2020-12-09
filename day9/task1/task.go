package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var windowSize = 25

func main() {
	dat, err := ioutil.ReadFile("./day9/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var numbers []int

	for _, row := range rows {
		n, err := strconv.Atoi(row)
		check(err)
		numbers = append(numbers, n)
	}

	lookup := make(map[int]int)

	for i := 0; i < windowSize; i++ {
		for j := i + 1; j < windowSize; j++ {
			s := numbers[i] + numbers[j]
			if c, ok := lookup[s]; ok {
				lookup[s] = c + 1
			} else {
				lookup[s] = 1
			}
		}
	}

	i := windowSize

	for i < len(numbers) {
		// fmt.Printf("Checking for %d in %v\n", numbers[i], lookup)
		if c, ok := lookup[numbers[i]]; !ok || c == 0 {
			fmt.Printf("%d is not a sum of any of %d previous numbers\n", numbers[i], windowSize)
			break
		}

		for j := 1; j < windowSize; j++ {
			s := numbers[i-windowSize] + numbers[i-j]
			lookup[s] = lookup[s] - 1
		}

		for j := 1; j < windowSize; j++ {
			s := numbers[i] + numbers[i-j]
			if c, ok := lookup[s]; ok {
				lookup[s] = c + 1
			} else {
				lookup[s] = 1
			}
		}

		i++
	}

}
