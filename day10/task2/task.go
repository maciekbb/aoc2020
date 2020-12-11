package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day10/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var numbers []int

	numbers = append(numbers, 0)

	for _, row := range rows {
		voltage, err := strconv.Atoi(row)
		check(err)
		numbers = append(numbers, voltage)
	}

	sort.Ints(numbers)

	device := numbers[len(numbers)-1] + 3

	lookup := make(map[int]bool)

	for _, n := range numbers {
		lookup[n] = true
	}

	cache := make(map[int]int)
	var solve func(start int) int

	solve = func(voltage int) int {
		if device == voltage {
			return 1
		}

		if v, ok := cache[voltage]; ok {
			return v
		}

		if _, ok := lookup[voltage]; !ok {
			return 0
		}

		r := solve(voltage+1) + solve(voltage+2) + solve(voltage+3)
		cache[voltage] = r
		return r
	}

	r := solve(0)
	fmt.Printf("Soluton is %d\n", r)
}
